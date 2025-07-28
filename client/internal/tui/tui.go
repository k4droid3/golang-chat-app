package tui

import (
	"github.com/k4droid3/TUI-chat/internal/chat"
)

type UI struct {
	State *chat.State
}

func NewTui() *UI {
	return &UI{}
}

func (ui *UI) Run() {

	state := chat.State{
		User:    "Ciri",
		Status:  chat.Online,
		History: []chat.Message{},
	}

	ui.State = &state

}
