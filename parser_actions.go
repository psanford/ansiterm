package ansiterm

func (ap *AnsiParser) collectParam() error {
	currChar := ap.context.CurrentChar()
	ap.logf("collectParam %#x", currChar)
	ap.context.CurCharToParamBuffer()
	return nil
}

func (ap *AnsiParser) collectInter() error {
	// TODO Implement DCS intermediate state handling
	return nil
}

func (ap *AnsiParser) escDispatch() error {
	cmd, _ := parseCmd(*ap.context)
	ap.logf("escDispatch currentChar: %#x", ap.context.CurrentChar())

	switch cmd {
	case "D": // IND
		// return ap.eventHandler.IND()
		e := &Index{
			raw: ap.context.Raw(),
		}
		ap.emit(e)
	case "E": // NEL, equivalent to CRLF
		e := &Execute{
			raw: ap.context.Raw(),
			B:   []byte("\r\n"),
		}
		ap.emit(e)
	case "M": // RI
		e := &ReverseIndex{
			raw: ap.context.Raw(),
		}
		ap.emit(e)
	}

	return nil
}

func (ap *AnsiParser) csiDispatch() error {
	cmd, _ := parseCmd(*ap.context)
	params, _ := parseParams(ap.context.ParamBuffer())
	ap.logf("Parsed params: %v with length: %d", params, len(params))

	ap.logf("csiDispatch: %v(%v)", cmd, params)

	switch cmd {
	case "@":
		// return ap.eventHandler.ICH(getInt(params, 1))
		e := &InsertCharacter{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "A":
		// return ap.eventHandler.CUU(getInt(params, 1))
		e := &CursorUp{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "B":
		// return ap.eventHandler.CUD(getInt(params, 1))
		e := &CursorDown{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "C":
		// return ap.eventHandler.CUF(getInt(params, 1))
		e := &CursorForward{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "D":
		// return ap.eventHandler.CUB(getInt(params, 1))
		e := &CursorBackward{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "E":
		// return ap.eventHandler.CNL(getInt(params, 1))
		e := &CursorNextLine{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "F":
		// return ap.eventHandler.CPL(getInt(params, 1))
		e := &CursorPreviousLine{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "G":
		// return ap.eventHandler.CHA(getInt(params, 1))
		e := &CursorHorizontalAbsolute{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "H":
		ints := getInts(params, 2, 1)
		x, y := ints[0], ints[1]
		// return ap.eventHandler.CUP(x, y)
		e := &CursorPosition{
			raw: ap.context.Raw(),
			Row: y,
			Col: x,
		}
		ap.emit(e)
	case "J":
		// param := getEraseParam(params)
		// return ap.eventHandler.ED(param)
		e := &EraseInDisplay{
			raw: ap.context.Raw(),
			N:   getEraseParam(params),
		}
		ap.emit(e)
	case "K":
		// param := getEraseParam(params)
		// return ap.eventHandler.EL(param)
		e := &EraseInLine{
			raw: ap.context.Raw(),
			N:   getEraseParam(params),
		}
		ap.emit(e)
	case "L":
		// return ap.eventHandler.IL(getInt(params, 1))
		e := &InsertLine{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "M":
		// return ap.eventHandler.DL(getInt(params, 1))
		e := &DeleteLine{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "P":
		// return ap.eventHandler.DCH(getInt(params, 1))
		e := &DeleteCharacter{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "S":
		// return ap.eventHandler.SU(getInt(params, 1))
		e := &ScrollUp{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "T":
		// return ap.eventHandler.SD(getInt(params, 1))
		e := &ScrollDown{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "c":
		// return ap.eventHandler.DA(params)
		e := &DeviceAttributes{
			raw:        ap.context.Raw(),
			Attributes: params,
		}
		ap.emit(e)
	case "d":
		// return ap.eventHandler.VPA(getInt(params, 1))
		e := &VerticalLinePositionAbsolute{
			raw: ap.context.Raw(),
			N:   getInt(params, 1),
		}
		ap.emit(e)
	case "f":
		ints := getInts(params, 2, 1)
		x, y := ints[0], ints[1]
		// return ap.eventHandler.HVP(x, y)
		e := &HorizontalVerticalPosition{
			raw: ap.context.Raw(),
			Row: y,
			Col: x,
		}
		ap.emit(e)
	case "h":
		return ap.hDispatch(params)
	case "l":
		return ap.lDispatch(params)
	case "m":
		// return ap.eventHandler.SGR(getInts(params, 1, 0))
		e := &SetGraphicsRendition{
			raw:  ap.context.Raw(),
			Attr: getInts(params, 1, 0),
		}
		ap.emit(e)
	case "r":
		ints := getInts(params, 2, 1)
		// top, bottom := ints[0], ints[1]
		// return ap.eventHandler.DECSTBM(top, bottom)
		e := &SetTopAndBottomMargins{
			raw:    ap.context.Raw(),
			Top:    ints[0],
			Bottom: ints[1],
		}
		ap.emit(e)
	case "~":
		ints := getInts(params, 1, 0)
		if ints[0] == 3 {
			e := &DeleteCharacter{
				raw: ap.context.Raw(),
				N:   1,
			}
			ap.emit(e)
		} else {
			// Dispatch generic CSI
		}
	default:
		ap.logf("ERROR: Unsupported CSI command: '%s', with full context:  %v", cmd, ap.context)
		return nil
	}

	return nil
}

func (ap *AnsiParser) print() error {
	e := &Print{
		raw: ap.context.Raw(),
		B:   []byte{ap.context.CurrentChar()},
	}
	ap.emit(e)
	return nil
}

func (ap *AnsiParser) clear() error {
	ap.context.Reset()
	return nil
}

func (ap *AnsiParser) execute() error {
	e := &Execute{
		raw: ap.context.Raw(),
		B:   []byte{ap.context.CurrentChar()},
	}
	ap.emit(e)
	return nil
}

func (ap *AnsiParser) emit(e AnsiEvent) {
	ap.eventChan <- e
	ap.context.Reset()
}
