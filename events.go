package ansiterm

type AnsiEvent interface {
	Raw() []byte
}

// Print
type Print struct {
	raw []byte
	B   []byte
}

func (e *Print) Raw() []byte {
	return e.raw
}

// Execute C0 commands.
// This is ascii 0x00-1F, 0x7F
type Execute struct {
	raw []byte
	B   []byte
}

func (e *Execute) Raw() []byte {
	return e.raw
}

// CursorUp (CUU)
type CursorUp struct {
	raw []byte
	N   int // number of lines to move the cursor up by
}

func (e *CursorUp) Raw() []byte {
	return e.raw
}

// CursorDown (CUD)
type CursorDown struct {
	raw []byte
	N   int
}

func (e *CursorDown) Raw() []byte {
	return e.raw
}

// CursorForward (CUF)
type CursorForward struct {
	raw []byte
	N   int
}

func (e *CursorForward) Raw() []byte {
	return e.raw
}

// CursorBackward (CUB)
type CursorBackward struct {
	raw []byte
	N   int
}

func (e *CursorBackward) Raw() []byte {
	return e.raw
}

// CursorNextLine (CNL)
type CursorNextLine struct {
	raw []byte
	N   int
}

func (e *CursorNextLine) Raw() []byte {
	return e.raw
}

// CursorPreviousLine (CPL)
type CursorPreviousLine struct {
	raw []byte
	N   int
}

func (e *CursorPreviousLine) Raw() []byte {
	return e.raw
}

// CursorHorizontalAbsolute (CHA)
type CursorHorizontalAbsolute struct {
	raw []byte
	N   int
}

func (e *CursorHorizontalAbsolute) Raw() []byte {
	return e.raw
}

// VerticalLinePositionAbsolute (VPA)
type VerticalLinePositionAbsolute struct {
	raw []byte
	N   int
}

func (e *VerticalLinePositionAbsolute) Raw() []byte {
	return e.raw
}

// CursorPosition (CUP)
type CursorPosition struct {
	raw []byte
	Row int
	Col int
}

func (e *CursorPosition) Raw() []byte {
	return e.raw
}

// HorizontalVerticalPosition (HVP)
type HorizontalVerticalPosition struct {
	raw []byte
	Row int
	Col int
}

func (e *HorizontalVerticalPosition) Raw() []byte {
	return e.raw
}

// TextCursorEnableMode (DECTCEM)
type TextCursorEnableMode struct {
	raw    []byte
	Enable bool
}

func (e *TextCursorEnableMode) Raw() []byte {
	return e.raw
}

// OriginMode (DECOM)
type OriginMode struct {
	raw    []byte
	Enable bool
}

func (e *OriginMode) Raw() []byte {
	return e.raw
}

// ColumnMode (DECCOLM)
type ColumnMode struct {
	raw    []byte
	Enable bool
}

func (e *ColumnMode) Raw() []byte {
	return e.raw
}

// EraseInDisplay (ED)
type EraseInDisplay struct {
	raw []byte
	N   int
}

func (e *EraseInDisplay) Raw() []byte {
	return e.raw
}

// EraseInLine (EL)
type EraseInLine struct {
	raw []byte
	N   int
}

func (e *EraseInLine) Raw() []byte {
	return e.raw
}

// InsertLine (IL)
type InsertLine struct {
	raw []byte
	N   int
}

func (e *InsertLine) Raw() []byte {
	return e.raw
}

// DeleteLine (DL)
type DeleteLine struct {
	raw []byte
	N   int
}

func (e *DeleteLine) Raw() []byte {
	return e.raw
}

// InsertCharacter (ICH)
type InsertCharacter struct {
	raw []byte
	N   int
}

func (e *InsertCharacter) Raw() []byte {
	return e.raw
}

// DeleteCharacter (DCH)
type DeleteCharacter struct {
	raw []byte
	N   int
}

func (e *DeleteCharacter) Raw() []byte {
	return e.raw
}

// SetGraphicsRendition (SGR)
type SetGraphicsRendition struct {
	raw  []byte
	Attr []int
}

func (e *SetGraphicsRendition) Raw() []byte {
	return e.raw
}

// ScrollUp (SU)
type ScrollUp struct {
	raw []byte
	N   int
}

func (e *ScrollUp) Raw() []byte {
	return e.raw
}

// ScrollDown (SD)
type ScrollDown struct {
	raw []byte
	N   int
}

func (e *ScrollDown) Raw() []byte {
	return e.raw
}

// DeviceAttributes (DA)
type DeviceAttributes struct {
	raw        []byte
	Attributes []string
}

func (e *DeviceAttributes) Raw() []byte {
	return e.raw
}

// SetTopAndBottomMargins (DECSTBM)
type SetTopAndBottomMargins struct {
	raw    []byte
	Top    int
	Bottom int
}

func (e *SetTopAndBottomMargins) Raw() []byte {
	return e.raw
}

// Index (IND)
type Index struct {
	raw []byte
}

func (e *Index) Raw() []byte {
	return e.raw
}

// ReverseIndex (RI)
type ReverseIndex struct {
	raw []byte
}

func (e *ReverseIndex) Raw() []byte {
	return e.raw
}

// generic in an event that we do not have a more specific state for
type generic struct {
	raw []byte
}

func (e *generic) Raw() []byte {
	return e.raw
}
