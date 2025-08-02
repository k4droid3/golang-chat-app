package client

import (
	"fmt"
	"os"
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
		Conn: chat.NewConnHandler("127.0.0.1:8080"),
		Term: term.NewTermHandler(),
	}
}

func (c *Client) Stop() {
	c.Conn.Disconnect()
	c.Term.Stop()
}

func (c *Client) Start() error {
	if err := c.Term.Start(); err != nil {
		return err
	}
	if err := c.Conn.Connect(); err != nil {
		return err
	}
	defer c.Stop()

	// // Mock Messages
	// c.Ui.View.History = []chat.Message{{User: "Geralt", Content: "Hmm.... a round of Gwent?", Timestamp: time.Now().Add(-20 * time.Second)},
	// 	{User: "Yennefer", Content: "We are at a funeral!", Timestamp: time.Now()}}

	go func() {
		<-c.Term.InterruptSignal
		fmt.Println("Received Interrupt Signal. Stopping...")
		c.Stop()
		os.Exit(1)
	}()

	go func() {
		for msg := range c.Conn.Recieve {
			c.Ui.View.History = append(c.Ui.View.History, msg)
			c.Ui.Render()
		}
	}()

	c.Ui.Render()
	for {
		inputChar := <-c.Term.BufChar
		if inputChar == '\n' {
			if c.Ui.View.InputBuffer == "/exit" {
				break
			}
			c.Conn.Send(c.Ui.View.InputBuffer)
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
