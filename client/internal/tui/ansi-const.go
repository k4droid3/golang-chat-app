package tui

import (
	"fmt"
)

// ANSI Escape Codes
const (
	// Clear Screen
	Clear = "\033[2J"

	// Reset all attributes
	Reset = "\033[0m"

	// Text Colors
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
)

func moveCursor(x, y int) string {
	return fmt.Sprintf("\033[%d;%dH", x, y)
}
