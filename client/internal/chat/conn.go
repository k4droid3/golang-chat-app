package chat

import (
	"net"
	"time"
)

type ConnHandler struct {
	Socket  string
	Conn    net.Conn
	Status  ConnStatus
	Recieve chan Message
}

func NewConnHandler(socket string) *ConnHandler {
	return &ConnHandler{
		Socket:  socket,
		Status:  Offline,
		Recieve: make(chan Message),
	}
}

func (c *ConnHandler) Connect() error {
	conn, err := net.Dial("tcp", c.Socket)
	if err != nil {
		return err
	}
	c.Conn = conn
	c.Status = Online

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := c.Conn.Read(buf)
			if err != nil {
				c.Status = Offline
				return
			}
			if n > 0 {
				message := Message{
					Content:   string(buf[:n]),
					User:      "Server",
					Timestamp: time.Now(),
				}
				c.Recieve <- message
			}
		}
	}()

	return nil
}

func (c *ConnHandler) Send(msg string) error {
	_, err := c.Conn.Write([]byte(msg + "\n"))
	return err
}
