package client

import (
	"time"

	"github.com/k4droid3/TUI-chat/internal/chat"
	"github.com/k4droid3/TUI-chat/internal/term"
	"github.com/k4droid3/TUI-chat/internal/tui"
)

type Client struct {
	Ui   *tui.Tui
	Conn *chat.ConnHandler
	Term *term.TermHandler
}

func NewClient(user string) *Client {
	return &Client{
		Ui: tui.NewTui(user, 24, 80),
	}
}

func (c *Client) Start() error {
	for {
		c.Ui.Render()
		time.Sleep(100 * time.Millisecond)
	}
}
