package chat

import (
	"fmt"
	"net"
	"time"
)

type ClientHandler struct {
	User       *User
	Connection net.Conn
	Send       chan Message
	Room       *Room
}

func (handler *ClientHandler) Run() {
	// defer handler.Connection.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := handler.Connection.Read(buf)
			if err != nil {
				fmt.Println("[", handler.User.Username, "]", "Error reading data: ", err)
				handler.Room.Leave <- handler
				return
			}
			message := Message{
				Content:   string(buf[:n]),
				Sender:    handler.User,
				CreatedAt: time.Now(),
			}
			handler.Room.Broadcast <- message
		}
	}()

	for msg := range handler.Send {
		handler.Connection.Write([]byte(fmt.Sprint(msg.Sender.Username, ":", msg.Content)))
	}
}
