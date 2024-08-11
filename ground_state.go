package ansiterm

type groundState struct {
	baseState
}

func (gs groundState) Handle(b byte) (s state, e error) {
	gs.parser.logf("AnsiParser Handle char %#02x", b)
	nextState, err := gs.baseState.Handle(b)
	if err != nil {
		return nextState, err
	}
	if nextState != nil && nextState != gs {
		gs.parser.logf("AnsiParser nope early %#02x, %s %s", b, nextState.Name(), err)
		return nextState, err
	}

	switch {
	case sliceContains(printables, b):
		gs.parser.logf("AnsiParser print char %#02x", b)
		return gs, gs.parser.print()
	case sliceContains(executors, b):
		gs.parser.logf("AnsiParser execute char %#02x", b)
		return gs, gs.parser.execute()
	}

	if b&0x80 == 0x80 { // the msb is set, assume this is a unicode or extend charset
		return gs.parser.utf8State, nil
	}

	return gs, nil
}
