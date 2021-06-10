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

	// Background colors
	BlackBG, RedBG, GreenBG, YellowBG, BlueBG, MagentaBG, CyanBG, WhiteBG []byte

	XTerm XTermColour

	// Reset all attributes
	Reset      []byte
	Clear      []byte // Clear screen.
	Clear0     []byte // Clear to end.
	Clear1     []byte // Clear to start.
	Clear2     []byte // Clear entire screen
	ClearLine  []byte // Clear line.
	ClearLine0 []byte // Clear line to end.
	ClearLine1 []byte // Clear line to start
	ClearLine2 []byte // Clear entire line.
	NextLine   []byte
	LLCorner   []byte
	Underline  []byte
	Reverse    []byte
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

	BlackBG:   []byte{keyEscape, '[', '4', '0', 'm'},
	RedBG:     []byte{keyEscape, '[', '4', '1', 'm'},
	GreenBG:   []byte{keyEscape, '[', '4', '2', 'm'},
	YellowBG:  []byte{keyEscape, '[', '4', '3', 'm'},
	BlueBG:    []byte{keyEscape, '[', '4', '4', 'm'},
	MagentaBG: []byte{keyEscape, '[', '4', '5', 'm'},
	CyanBG:    []byte{keyEscape, '[', '4', '6', 'm'},
	WhiteBG:   []byte{keyEscape, '[', '4', '7', 'm'},

	XTerm: Colours,

	Reset: []byte{keyEscape, '[', '0', 'm'},

	Clear:  []byte{keyEscape, '[', 'J'},      // Clear screen.
	Clear0: []byte{keyEscape, '[', '0', 'J'}, // Clear to end.
	Clear1: []byte{keyEscape, '[', '1', 'J'}, // Clear to start.
	Clear2: []byte{keyEscape, '[', '2', 'J'}, // Clear entire screen.

	ClearLine:  []byte{keyEscape, '[', 'J'},      // Clear line.
	ClearLine0: []byte{keyEscape, '[', '0', 'J'}, // Clear line to end.
	ClearLine1: []byte{keyEscape, '[', '1', 'J'}, // Clear line to start.
	ClearLine2: []byte{keyEscape, '[', '2', 'J'}, // Clear entire line.

	NextLine:  []byte{keyEscape, 'E'},
	LLCorner:  []byte{keyEscape, 'F'},
	Underline: []byte{keyEscape, '[', '4', 'm'},
	Reverse:   []byte{keyEscape, '[', '7', 'm'},
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
	Esc *EscapeCodes

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

// NewTerminal returns a terminal with the given prompt.
func NewTerminal(c io.ReadWriter, prompt string) *Terminal {
	return &Terminal{
		Esc:          &vt100EscapeCodes,
		c:            c,
		prompt:       []rune(prompt),
		termWidth:    80,
		termHeight:   24,
		echo:         true,
		historyIndex: -1,
	}
}
