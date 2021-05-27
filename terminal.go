package term

import (
	"io"
	"sync"
)

// EscapeCodes contains escape sequences that can be written to the
// terminal in order to achieve different styles of text.
type EscapeCodes struct {
	// Foreground colors
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White []byte

	XTerm XTermColour

	// Reset all attributes
	Reset    []byte
	Clear    []byte
	NextLine []byte
	LLCorner []byte
}

var vt100EscapeCodes = EscapeCodes{

	Black:   []byte{keyEscape, '[', '3', '0', 'm'},
	Red:     []byte{keyEscape, '[', '3', '1', 'm'},
	Green:   []byte{keyEscape, '[', '3', '2', 'm'},
	Yellow:  []byte{keyEscape, '[', '3', '3', 'm'},
	Blue:    []byte{keyEscape, '[', '3', '4', 'm'},
	Magenta: []byte{keyEscape, '[', '3', '5', 'm'},
	Cyan:    []byte{keyEscape, '[', '3', '6', 'm'},
	White:   []byte{keyEscape, '[', '3', '7', 'm'},

	XTerm: Colours,

	Reset:    []byte{keyEscape, '[', '0', 'm'},
	Clear:    []byte{keyEscape, '[', '2', 'J'},
	NextLine: []byte{keyEscape, 'E'},
	LLCorner: []byte{keyEscape, 'F'},
}

// Terminal contains the state for running a VT100 terminal that is capable of
// reading lines of input.
type Terminal struct {
	// AutoCompleteCallback, if non-null, is called for each keypress with
	// the full input line and the current position of the cursor (in
	// bytes, as an index into |line|). If it returns ok=false, the key
	// press is processed normally. Otherwise it returns a replacement line
	// and the new cursor position.
	AutoCompleteCallback func(line string, pos int, key rune) (newLine string, newPos int, ok bool)

	// Escape contains a pointer to the escape codes for this terminal.
	// It's always a valid pointer, although the escape codes themselves
	// may be empty if the terminal doesn't support them.
	Escape *EscapeCodes

	// lock protects the terminal and the state in this object from
	// concurrent processing of a key press and a Write() call.
	lock sync.Mutex

	c      io.ReadWriter
	prompt []rune

	// line is the current line being entered.
	line []rune
	// pos is the logical position of the cursor in line
	pos int
	// echo is true if local echo is enabled
	echo bool
	// pasteActive is true iff there is a bracketed paste operation in
	// progress.
	pasteActive bool

	// cursorX contains the current X value of the cursor where the left
	// edge is 0. cursorY contains the row number where the first row of
	// the current line is 0.
	cursorX, cursorY int
	// maxLine is the greatest value of cursorY so far.
	maxLine int

	termWidth, termHeight int

	// outBuf contains the terminal data to be sent.
	outBuf []byte
	// remainder contains the remainder of any partial key sequences after
	// a read. It aliases into inBuf.
	remainder []byte
	inBuf     [256]byte

	// history contains previously entered commands so that they can be
	// accessed with the up and down keys.
	history stRingBuffer
	// historyIndex stores the currently accessed history entry, where zero
	// means the immediately previous entry.
	historyIndex int
	// When navigating up and down the history it's possible to return to
	// the incomplete, initial line. That value is stored in
	// historyPending.
	historyPending string
}

func NewTerminal(c io.ReadWriter, prompt string) *Terminal {
	return &Terminal{
		Escape:       &vt100EscapeCodes,
		c:            c,
		prompt:       []rune(prompt),
		termWidth:    80,
		termHeight:   24,
		echo:         true,
		historyIndex: -1,
	}
}
