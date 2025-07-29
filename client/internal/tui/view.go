package tui

import (
	"github.com/k4droid3/TUI-chat/internal/chat"
)

type view struct {
	User        string
	Status      chat.ConnStatus
	History     []chat.Message
	InputBuffer string
	SentBuffer  []string
	Height      int
	Width       int
	//TODO: add lock
}

// should it have the same state as chat.go for view?
// it should definetly have input shown as well ig?
// maybe add height and all here as well
