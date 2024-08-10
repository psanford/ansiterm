package ansiterm

type ansiContext struct {
	cCurrentChar byte
	cParamBuffer []byte

	collectedBytes []byte
}

func (c *ansiContext) Reset() {
	c.cCurrentChar = 0
	c.cParamBuffer = []byte{}
	c.collectedBytes = []byte{}
}

func (c *ansiContext) CollectCurrentChar(b byte) {
	c.cCurrentChar = b
	c.collectedBytes = append(c.collectedBytes, b)
}

func (c *ansiContext) Raw() []byte {
	return c.collectedBytes
}

func (c *ansiContext) CurrentChar() byte {
	return c.cCurrentChar
}

func (c *ansiContext) CurCharToParamBuffer() {
	c.cParamBuffer = append(c.cParamBuffer, c.cCurrentChar)
}

func (c *ansiContext) ParamBuffer() []byte {
	return c.cParamBuffer
}
