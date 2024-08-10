package ansiterm

type ansiContext struct {
	currentChar byte
	paramBuffer []byte
}
