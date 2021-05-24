package term

import "golang.org/x/term"

// Terminal contains the state for running a VT100 terminal that is capable of
// reading lines of input.
type Terminal = term.Terminal

// EscapeCodes contains escape sequences that can be written to the terminal in
// order to achieve different styles of text.
type EscapeCodes = term.EscapeCodes

// State contains the state of a terminal.
type State = term.State

// IsTerminal returns whether the given file descriptor is a terminal.
func IsTerminal(fd int) bool {
	return term.IsTerminal(fd)
}

// MakeRaw puts the terminal connected to the given file descriptor into raw
// mode and returns the previous state of the terminal so that it can be
// restored.
func MakeRaw(fd int) (*State, error) {
	return term.MakeRaw(fd)
}

// GetState returns the current state of a terminal which may be useful to
// restore the terminal after a signal.
func GetState(fd int) (*State, error) {
	return term.GetState(fd)
}

// Restore restores the terminal connected to the given file descriptor to a
// previous state.
func Restore(fd int, oldState *State) error {
	return term.Restore(fd, oldState)
}

// GetSize returns the visible dimensions of the given terminal.
//
// These dimensions don't include any scrollback buffer height.
func GetSize(fd int) (width, height int, err error) {
	return term.GetSize(fd)
}

// ReadPassword reads a line of input from a terminal without local echo.  This
// is commonly used for inputting passwords and other sensitive data. The slice
// returned does not include the \n.
func ReadPassword(fd int) ([]byte, error) {
	return term.ReadPassword(fd)
}
