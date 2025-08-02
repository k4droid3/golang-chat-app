package chat

import (
	"net"
	"time"
)

type ConnHandler struct {
	socket  string
	conn    net.Conn
	Status  ConnStatus
	Recieve chan Message
}

func NewConnHandler(socket string) *ConnHandler {
	return &ConnHandler{
		socket:  socket,
		Status:  Offline,
		Recieve: make(chan Message),
	}
}

func (c *ConnHandler) Connect() error {
	conn, err := net.Dial("tcp", c.socket)
	if err != nil {
		return err
	}
	c.conn = conn
	c.Status = Online

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := c.conn.Read(buf)
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

func (c *ConnHandler) Disconnect() error {
	return c.conn.Close()
}

func (c *ConnHandler) Send(msg string) error {
	_, err := c.conn.Write([]byte(msg + "\n"))
	return err
}
