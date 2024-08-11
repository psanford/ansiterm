package ansiterm

type utf8State struct {
	baseState
}

func (s utf8State) Handle(b byte) (state, error) {
	s.parser.logf("AnsiParse::utf8 Handle byte %#02x", b)

	buf := s.parser.context.cParamBuffer
	buf = append(buf, s.parser.context.cCurrentChar)

	switch utf8StateCheck(buf) {
	case utf8Invalid:
		s.parser.logf("AnsiParser::utf8 invalid %#02x", b)
		return s.parser.ground, s.parser.print()
	case utf8ValidPartial:
		s.parser.logf("AnsiParser::utf8 valid partial")
		s.parser.context.CurCharToParamBuffer()
		return s, nil
	case utf8ValidFull:
		s.parser.logf("AnsiParser::utf8 valid full")
		return s.parser.ground, s.parser.print()
	}

	return s.parser.ground, nil
}

func (s utf8State) Enter() error {
	s.parser.context.CurCharToParamBuffer()
	return nil
}

type utf8ParserState int

const (
	utf8Invalid      utf8ParserState = 0
	utf8ValidPartial utf8ParserState = iota + 1
	utf8ValidFull
)

func utf8StateCheck(c []byte) utf8ParserState {
	if len(c) == 0 {
		return utf8Invalid
	}
	if len(c) == 1 {
		if c[0]&0x80 == 0x00 {
			return utf8ValidFull
		}
	}

	expectLen := -1
	first := c[0]

	if first>>3 == 0x1e { // 11110b
		expectLen = 4
	} else if first>>4 == 0x0e { // 1110b
		expectLen = 3
	} else if first>>5 == 0x06 { // 110b
		expectLen = 2
	} else {
		return utf8Invalid
	}

	if len(c) < expectLen {
		return utf8ValidPartial
	} else if len(c) > expectLen {
		return utf8Invalid
	}

	for i := 1; i < len(c); i++ {
		b := c[i]
		if b>>6 != 0x02 { // 10b
			return utf8Invalid
		}
	}

	return utf8ValidFull
}
