package ansiterm

import (
	"strconv"
)

func parseParams(bytes []byte) ([]string, error) {
	paramBuff := make([]byte, 0, 0)
	params := []string{}

	for _, v := range bytes {
		if v == ';' {
			if len(paramBuff) > 0 {
				// Completed parameter, append it to the list
				s := string(paramBuff)
				params = append(params, s)
				paramBuff = make([]byte, 0, 0)
			}
		} else {
			paramBuff = append(paramBuff, v)
		}
	}

	// Last parameter may not be terminated with ';'
	if len(paramBuff) > 0 {
		s := string(paramBuff)
		params = append(params, s)
	}

	return params, nil
}

func parseCmd(context ansiContext) (string, error) {
	return string(context.CurrentChar()), nil
}

func getInt(params []string, dflt int) int {
	i := getInts(params, 1, dflt)[0]
	return i
}

func getInts(params []string, minCount int, dflt int) []int {
	ints := []int{}

	for _, v := range params {
		i, _ := strconv.Atoi(v)
		// Zero is mapped to the default value in VT100.
		if i == 0 {
			i = dflt
		}
		ints = append(ints, i)
	}

	if len(ints) < minCount {
		remaining := minCount - len(ints)
		for i := 0; i < remaining; i++ {
			ints = append(ints, dflt)
		}
	}

	return ints
}

func (ap *AnsiParser) modeDispatch(param string, set bool) error {
	switch param {
	case "?3":
		e := &ColumnMode{
			raw:    ap.context.Raw(),
			Enable: set,
		}
		ap.eventChan <- e
	case "?6":
		e := &OriginMode{
			raw:    ap.context.Raw(),
			Enable: set,
		}
		ap.eventChan <- e
	case "?25":
		e := &TextCursorEnableMode{
			raw:    ap.context.Raw(),
			Enable: set,
		}
		ap.eventChan <- e
	}
	return nil
}

func (ap *AnsiParser) hDispatch(params []string) error {
	if len(params) == 1 {
		return ap.modeDispatch(params[0], true)
	}

	return nil
}

func (ap *AnsiParser) lDispatch(params []string) error {
	if len(params) == 1 {
		return ap.modeDispatch(params[0], false)
	}

	return nil
}

func getEraseParam(params []string) int {
	param := getInt(params, 0)
	if param < 0 || 3 < param {
		param = 0
	}

	return param
}
