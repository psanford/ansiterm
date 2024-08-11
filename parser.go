package ansiterm

import (
	"context"
	"errors"
)

type AnsiParser struct {
	currState          state
	eventHandler       ansiEventHandler
	eventChan          chan<- AnsiEvent
	context            *ansiContext
	csiEntry           state
	csiParam           state
	dcsEntry           state
	escape             state
	escapeIntermediate state
	error              state
	ground             state
	oscString          state
	utf8State          state
	stateMap           []state

	initialState string
	logf         func(string, ...interface{})
}

type Option func(*AnsiParser)

func WithLogf(f func(string, ...interface{})) Option {
	return func(ap *AnsiParser) {
		ap.logf = f
	}
}

func WithInitialState(s string) Option {
	return func(ap *AnsiParser) {
		ap.initialState = s
	}
}

func createParserEventHandler(ctx context.Context, evtHandler ansiEventHandler, opts ...Option) *AnsiParser {
	eventChan := make(chan AnsiEvent)
	ap := CreateParser(eventChan, opts...)
	go func() {
		for {
			var evt AnsiEvent
			select {
			case evt = <-eventChan:
			case <-ctx.Done():
				return
			}

			switch e := evt.(type) {
			case *Print:
				evtHandler.Print(e.B)
			case *Execute:
				evtHandler.Execute(e.B[0])
			case *CursorUp:
				evtHandler.CUU(e.N)
			case *CursorDown:
				evtHandler.CUD(e.N)
			case *CursorForward:
				evtHandler.CUF(e.N)
			case *CursorBackward:
				evtHandler.CUB(e.N)
			case *CursorNextLine:
				evtHandler.CNL(e.N)
			case *CursorPreviousLine:
				evtHandler.CPL(e.N)
			case *CursorHorizontalAbsolute:
				evtHandler.CHA(e.N)
			case *VerticalLinePositionAbsolute:
				evtHandler.VPA(e.N)
			case *CursorPosition:
				evtHandler.CUP(e.Col, e.Row)
			case *HorizontalVerticalPosition:
				evtHandler.HVP(e.Col, e.Row)
			case *TextCursorEnableMode:
				evtHandler.DECTCEM(e.Enable)
			case *OriginMode:
				evtHandler.DECOM(e.Enable)
			case *ColumnMode:
				evtHandler.DECCOLM(e.Enable)
			case *EraseInDisplay:
				evtHandler.ED(e.N)
			case *EraseInLine:
				evtHandler.EL(e.N)
			case *InsertLine:
				evtHandler.IL(e.N)
			case *DeleteLine:
				evtHandler.DL(e.N)
			case *InsertCharacter:
				evtHandler.ICH(e.N)
			case *DeleteCharacter:
				evtHandler.DCH(e.N)
			case *SetGraphicsRendition:
				evtHandler.SGR(e.Attr)
			case *ScrollUp:
				evtHandler.SU(e.N)
			case *ScrollDown:
				evtHandler.SD(e.N)
			case *DeviceAttributes:
				evtHandler.DA(e.Attributes)
			case *SetTopAndBottomMargins:
				evtHandler.DECSTBM(e.Top, e.Bottom)
			case *Index:
				evtHandler.IND()
			case *ReverseIndex:
				evtHandler.RI()
			default:
				// Handle unknown event types
				ap.logf("Unknown event type: %T", e)
			}

			evtHandler.Flush()
		}
	}()

	return ap
}

func CreateParser(eventChan chan<- AnsiEvent, opts ...Option) *AnsiParser {
	ap := &AnsiParser{
		eventChan:    eventChan,
		context:      &ansiContext{},
		initialState: "Ground",
	}
	for _, o := range opts {
		o(ap)
	}

	if ap.logf == nil {
		ap.logf = func(string, ...interface{}) {}
	}

	ap.csiEntry = csiEntryState{baseState{name: "CsiEntry", parser: ap}}
	ap.csiParam = csiParamState{baseState{name: "CsiParam", parser: ap}}
	ap.dcsEntry = dcsEntryState{baseState{name: "DcsEntry", parser: ap}}
	ap.escape = escapeState{baseState{name: "Escape", parser: ap}}
	ap.escapeIntermediate = escapeIntermediateState{baseState{name: "EscapeIntermediate", parser: ap}}
	ap.error = errorState{baseState{name: "Error", parser: ap}}
	ap.ground = groundState{baseState{name: "Ground", parser: ap}}
	ap.oscString = oscStringState{baseState{name: "OscString", parser: ap}}
	ap.utf8State = utf8State{baseState{name: "Utf8", parser: ap}}

	ap.stateMap = []state{
		ap.csiEntry,
		ap.csiParam,
		ap.dcsEntry,
		ap.escape,
		ap.escapeIntermediate,
		ap.error,
		ap.ground,
		ap.oscString,
		ap.utf8State,
	}

	ap.currState = getState(ap.initialState, ap.stateMap)

	ap.logf("CreateParser: parser %p", ap)
	return ap
}

func getState(name string, states []state) state {
	for _, el := range states {
		if el.Name() == name {
			return el
		}
	}

	return nil
}

func (ap *AnsiParser) Parse(bytes []byte) (int, error) {
	for i, b := range bytes {
		if err := ap.handle(b); err != nil {
			return i, err
		}
	}

	return len(bytes), nil
}

func (ap *AnsiParser) handle(b byte) error {
	ap.logf("AnsiParser handle: <%#02x> curstate=%s", b, ap.currState.Name())
	ap.context.CollectCurrentChar(b)
	newState, err := ap.currState.Handle(b)
	if err != nil {
		ap.logf("currState %s handle err: %s", ap.currState.Name(), err)
		return err
	}

	if newState == nil {
		ap.logf("WARNING: newState is nil")
		return errors.New("New state of 'nil' is invalid.")
	}
	ap.logf("newstate %s", newState.Name())

	if newState != ap.currState {
		if err := ap.changeState(newState); err != nil {
			return err
		}
	}

	return nil
}

func (ap *AnsiParser) changeState(newState state) error {
	ap.logf("ChangeState %s --> %s", ap.currState.Name(), newState.Name())

	// Exit old state
	if err := ap.currState.Exit(); err != nil {
		ap.logf("Exit state '%s' failed with : '%v'", ap.currState.Name(), err)
		return err
	}

	// Perform transition action
	if err := ap.currState.Transition(newState); err != nil {
		ap.logf("Transition from '%s' to '%s' failed with: '%v'", ap.currState.Name(), newState.Name, err)
		return err
	}

	// Enter new state
	if err := newState.Enter(); err != nil {
		ap.logf("Enter state '%s' failed with: '%v'", newState.Name(), err)
		return err
	}

	ap.currState = newState
	return nil
}
