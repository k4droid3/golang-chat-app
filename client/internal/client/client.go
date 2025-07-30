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
		Ui:   tui.NewTui(user, 24, 80),
		Term: term.NewTermHandler(),
	}
}

func (c *Client) Start() error {
	if err := c.Term.Start(); err != nil {
		return err
	}
	defer c.Term.Stop()

	// Mock Messages
	c.Ui.View.History = []chat.Message{{User: "Geralt", Content: "Hmm.... a round of Gwent?", Timestamp: time.Now().Add(-20 * time.Second)},
		{User: "Yennefer", Content: "We are at a funeral!", Timestamp: time.Now()}}

	c.Ui.Render()
	for len(c.Term.InterruptSignal) == 0 {
		inputChar := <-c.Term.BufChar
		// switch inputChar {
		// case '\n':
		// 	if c.Ui.View.InputBuffer == "/exit" {
		// 		break
		// 	}
		// 	c.Ui.View.InputBuffer = ""
		// case '\b', 127:
		// 	if len(c.Ui.View.InputBuffer) > 0 {
		// 		c.Ui.View.InputBuffer = c.Ui.View.InputBuffer[:len(c.Ui.View.InputBuffer)]
		// 	}
		// default:
		// 	c.Ui.View.InputBuffer += string(inputChar)
		// }
		if inputChar == '\n' {
			if c.Ui.View.InputBuffer == "/exit" {
				break
			}
			c.Ui.View.InputBuffer = ""
		} else if inputChar == '\b' || inputChar == 127 {
			if len(c.Ui.View.InputBuffer) > 0 {
				c.Ui.View.InputBuffer = c.Ui.View.InputBuffer[:len(c.Ui.View.InputBuffer)-1]
			}
		} else {
			c.Ui.View.InputBuffer += string(inputChar)
		}
		c.Ui.Render()
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}
