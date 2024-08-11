package ansiterm

type AnsiEvent interface {
	Raw() []byte
	ParserState() ParserState
}

type ParserState string

const (
	StatePrint   ParserState = "print"
	StateExecute ParserState = "execute"
	StateCUU     ParserState = "cuu"     // CUrsor Up
	StateCUD     ParserState = "cud"     // CUrsor Down
	StateCUF     ParserState = "cuf"     // CUrsor Forward
	StateCUB     ParserState = "cub"     // CUrsor Backward
	StateCNL     ParserState = "cnl"     // Cursor to Next Line
	StateCPL     ParserState = "cpl"     // Cursor to Previous Line
	StateCHA     ParserState = "cha"     // Cursor Horizontal position Absolute
	StateVPA     ParserState = "vpa"     // Vertical line Position Absolute
	StateCUP     ParserState = "cup"     // CUrsor Position
	StateHVP     ParserState = "hvp"     // Horizontal and Vertical Position (depends on PUM)
	StateDECTCEM ParserState = "dectcem" // Text Cursor Enable Mode
	StateDECOM   ParserState = "decom"   // Origin Mode
	StateDECCOLM ParserState = "deccolm" // 132 Column Mode
	StateED      ParserState = "ed"      // Erase in Display
	StateEL      ParserState = "el"      // Erase in Line
	StateIL      ParserState = "il"      // Insert Line
	StateDL      ParserState = "dl"      // Delete Line
	StateICH     ParserState = "ich"     // Insert Character
	StateDCH     ParserState = "dch"     // Delete Character
	StateSGR     ParserState = "sgr"     // Set Graphics Rendition
	StateSU      ParserState = "su"      // Pan Down
	StateSD      ParserState = "sd"      // Pan Up
	StateDA      ParserState = "da"      // Device Attributes
	StateDECSTBM ParserState = "decstbm" // Set Top and Bottom Margins
	StateIND     ParserState = "ind"     // Index
	StateRI      ParserState = "ri"      // Reverse Index

	StateGeneric ParserState = "generic"
)

// Print
type Print struct {
	raw []byte
	B   []byte
}

func (e *Print) Raw() []byte {
	return e.raw
}

func (e *Print) ParserState() ParserState {
	return StatePrint
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

func (e *Execute) ParserState() ParserState {
	return StateExecute
}

// CursorUp (CUU)
type CursorUp struct {
	raw []byte
	N   int // number of lines to move the cursor up by
}

func (e *CursorUp) Raw() []byte {
	return e.raw
}

func (e *CursorUp) ParserState() ParserState {
	return StateCUU
}

// CursorDown (CUD)
type CursorDown struct {
	raw []byte
	N   int
}

func (e *CursorDown) Raw() []byte {
	return e.raw
}

func (e *CursorDown) ParserState() ParserState {
	return StateCUD
}

// CursorForward (CUF)
type CursorForward struct {
	raw []byte
	N   int
}

func (e *CursorForward) Raw() []byte {
	return e.raw
}

func (e *CursorForward) ParserState() ParserState {
	return StateCUF
}

// CursorBackward (CUB)
type CursorBackward struct {
	raw []byte
	N   int
}

func (e *CursorBackward) Raw() []byte {
	return e.raw
}

func (e *CursorBackward) ParserState() ParserState {
	return StateCUB
}

// CursorNextLine (CNL)
type CursorNextLine struct {
	raw []byte
	N   int
}

func (e *CursorNextLine) Raw() []byte {
	return e.raw
}

func (e *CursorNextLine) ParserState() ParserState {
	return StateCNL
}

// CursorPreviousLine (CPL)
type CursorPreviousLine struct {
	raw []byte
	N   int
}

func (e *CursorPreviousLine) Raw() []byte {
	return e.raw
}

func (e *CursorPreviousLine) ParserState() ParserState {
	return StateCPL
}

// CursorHorizontalAbsolute (CHA)
type CursorHorizontalAbsolute struct {
	raw []byte
	N   int
}

func (e *CursorHorizontalAbsolute) Raw() []byte {
	return e.raw
}

func (e *CursorHorizontalAbsolute) ParserState() ParserState {
	return StateCHA
}

// VerticalLinePositionAbsolute (VPA)
type VerticalLinePositionAbsolute struct {
	raw []byte
	N   int
}

func (e *VerticalLinePositionAbsolute) Raw() []byte {
	return e.raw
}

func (e *VerticalLinePositionAbsolute) ParserState() ParserState {
	return StateVPA
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

func (e *CursorPosition) ParserState() ParserState {
	return StateCUP
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

func (e *HorizontalVerticalPosition) ParserState() ParserState {
	return StateHVP
}

// TextCursorEnableMode (DECTCEM)
type TextCursorEnableMode struct {
	raw    []byte
	Enable bool
}

func (e *TextCursorEnableMode) Raw() []byte {
	return e.raw
}

func (e *TextCursorEnableMode) ParserState() ParserState {
	return StateDECTCEM
}

// OriginMode (DECOM)
type OriginMode struct {
	raw    []byte
	Enable bool
}

func (e *OriginMode) Raw() []byte {
	return e.raw
}

func (e *OriginMode) ParserState() ParserState {
	return StateDECOM
}

// ColumnMode (DECCOLM)
type ColumnMode struct {
	raw    []byte
	Enable bool
}

func (e *ColumnMode) Raw() []byte {
	return e.raw
}

func (e *ColumnMode) ParserState() ParserState {
	return StateDECCOLM
}

// EraseInDisplay (ED)
type EraseInDisplay struct {
	raw []byte
	N   int
}

func (e *EraseInDisplay) Raw() []byte {
	return e.raw
}

func (e *EraseInDisplay) ParserState() ParserState {
	return StateED
}

// EraseInLine (EL)
type EraseInLine struct {
	raw []byte
	N   int
}

func (e *EraseInLine) Raw() []byte {
	return e.raw
}

func (e *EraseInLine) ParserState() ParserState {
	return StateEL
}

// InsertLine (IL)
type InsertLine struct {
	raw []byte
	N   int
}

func (e *InsertLine) Raw() []byte {
	return e.raw
}

func (e *InsertLine) ParserState() ParserState {
	return StateIL
}

// DeleteLine (DL)
type DeleteLine struct {
	raw []byte
	N   int
}

func (e *DeleteLine) Raw() []byte {
	return e.raw
}

func (e *DeleteLine) ParserState() ParserState {
	return StateDL
}

// InsertCharacter (ICH)
type InsertCharacter struct {
	raw []byte
	N   int
}

func (e *InsertCharacter) Raw() []byte {
	return e.raw
}

func (e *InsertCharacter) ParserState() ParserState {
	return StateICH
}

// DeleteCharacter (DCH)
type DeleteCharacter struct {
	raw []byte
	N   int
}

func (e *DeleteCharacter) Raw() []byte {
	return e.raw
}

func (e *DeleteCharacter) ParserState() ParserState {
	return StateDCH
}

// SetGraphicsRendition (SGR)
type SetGraphicsRendition struct {
	raw  []byte
	Attr []int
}

func (e *SetGraphicsRendition) Raw() []byte {
	return e.raw
}

func (e *SetGraphicsRendition) ParserState() ParserState {
	return StateSGR
}

// ScrollUp (SU)
type ScrollUp struct {
	raw []byte
	N   int
}

func (e *ScrollUp) Raw() []byte {
	return e.raw
}

func (e *ScrollUp) ParserState() ParserState {
	return StateSU
}

// ScrollDown (SD)
type ScrollDown struct {
	raw []byte
	N   int
}

func (e *ScrollDown) Raw() []byte {
	return e.raw
}

func (e *ScrollDown) ParserState() ParserState {
	return StateSD
}

// DeviceAttributes (DA)
type DeviceAttributes struct {
	raw        []byte
	Attributes []string
}

func (e *DeviceAttributes) Raw() []byte {
	return e.raw
}

func (e *DeviceAttributes) ParserState() ParserState {
	return StateDA
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

func (e *SetTopAndBottomMargins) ParserState() ParserState {
	return StateDECSTBM
}

// Index (IND)
type Index struct {
	raw []byte
}

func (e *Index) Raw() []byte {
	return e.raw
}

func (e *Index) ParserState() ParserState {
	return StateIND
}

// ReverseIndex (RI)
type ReverseIndex struct {
	raw []byte
}

func (e *ReverseIndex) Raw() []byte {
	return e.raw
}

func (e *ReverseIndex) ParserState() ParserState {
	return StateRI
}

// generic in an event that we do not have a more specific state for
type generic struct {
	raw []byte
}

func (e *generic) Raw() []byte {
	return e.raw
}

func (e *generic) ParserState() ParserState {
	return StateGeneric
}
