package tui

import (
	"fmt"
	"strings"

	"github.com/k4droid3/TUI-chat/internal/chat"
)

type Tui struct {
	View *view
	Send chan chat.Message
	// Do we need a sentBuffer for message that were not sent due to conn issues?
}

func NewTui(user string, height int, width int) *Tui {
	return &Tui{
		View: &view{
			User:        user,
			Status:      chat.Offline,
			History:     []chat.Message{},
			InputBuffer: "",
			Height:      height,
			Width:       width,
		},
		Send: make(chan chat.Message),
	}
}

func (ui *Tui) Render() {
	// Clear Screen
	fmt.Print(Clear)

	// Seperator
	ui.seperator(1)

	// Status bar
	fmt.Print(moveCursor(2, 1), "|")
	fmt.Print(moveCursor(2, 2), "User: ", ui.View.User)
	fmt.Print(Red, moveCursor(2, ui.View.Width-len(ui.View.Status.String())), ui.View.Status.String(), Reset)
	fmt.Print(moveCursor(2, ui.View.Width), "|\n")

	// Seperator
	ui.seperator(3)

	// Chat history
	for i := 4; i < ui.View.Height-1; i++ {
		fmt.Print(moveCursor(i, 1), "|")
		if i-4 < len(ui.View.History) {
			messageDiv := fmt.Sprintf("%s: %s", ui.View.History[i-4].User, ui.View.History[i-4].Content)
			formattedTime := ui.View.History[i-4].Timestamp.Format("2006-01-02 15:04:05")
			fmt.Print(moveCursor(i, 2), messageDiv)
			fmt.Print(moveCursor(i, ui.View.Width-len(formattedTime)), formattedTime)
		}
		fmt.Print(moveCursor(i, ui.View.Width), "|\n")
	}

	// Seperator
	ui.seperator(ui.View.Height - 1)

	// Input Div
	fmt.Print(moveCursor(ui.View.Height, 0), "> ", ui.View.InputBuffer)
}

func (ui *Tui) seperator(row int) {
	fmt.Print(moveCursor(row, 1), "+")
	fmt.Print(moveCursor(row, 2), strings.Repeat("-", ui.View.Width-2))
	fmt.Print(moveCursor(row, ui.View.Width), "+\n")
}
