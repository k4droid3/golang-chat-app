package tui

import (
	"fmt"
	"strings"

	"github.com/k4droid3/TUI-chat/internal/chat"
)

type Tui struct {
	view *view
	send chan chat.Message
}

func NewTui(user string, height int, width int) *Tui {
	return &Tui{
		view: &view{
			User:        user,
			Status:      chat.Offline,
			History:     []chat.Message{},
			InputBuffer: "",
			SentBuffer:  []string{},
			Height:      height,
			Width:       width,
		},
		send: make(chan chat.Message),
	}
}

func (ui *Tui) Render() {
	// Clear Screen
	fmt.Print(Clear)

	// Seperator
	ui.seperator(1)

	// Status bar
	fmt.Print(moveCursor(2, 1), "|")
	fmt.Print(moveCursor(2, 2), ui.view.User)
	fmt.Print(Green + moveCursor(2, ui.view.Width-len(ui.view.Status.String())) + Reset)
	fmt.Print(moveCursor(2, ui.view.Width), "|\n")

	// Seperator
	ui.seperator(3)

	// Chat history
	for i := 4; i < ui.view.Height-1; i++ {
		fmt.Print(moveCursor(i, 1), "|")
		if i-4 < len(ui.view.History) {
			messageDiv := fmt.Sprintf("%s: %s", ui.view.History[i-4].User, ui.view.History[i-4].Content)
			formattedTime := ui.view.History[i-4].Timestamp.Format("2006-01-02 15:04:05")
			fmt.Print(moveCursor(i, 2), messageDiv)
			fmt.Print(moveCursor(i, ui.view.Width-len(formattedTime)), formattedTime)
		}
		fmt.Print(moveCursor(i, ui.view.Width), "|\n")
	}

	// Seperator
	ui.seperator(ui.view.Height - 1)

	// Input Div
	fmt.Print(moveCursor(ui.view.Height, 0), "> ", ui.view.InputBuffer)
}

func (ui *Tui) seperator(row int) {
	fmt.Print(moveCursor(row, 1), "+")
	fmt.Print(moveCursor(row, 2), strings.Repeat("-", ui.view.Width-2))
	fmt.Print(moveCursor(row, ui.view.Width), "+\n")
}
