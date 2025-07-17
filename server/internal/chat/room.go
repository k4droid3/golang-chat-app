package chat

import (
	"fmt"
)

type Room struct {
	Name      string
	Clients   map[*ClientHandler]bool
	Join      chan *ClientHandler
	Leave     chan *ClientHandler
	Broadcast chan Message
}

func NewRoom(name string) *Room {
	return &Room{
		Name:      name,
		Clients:   make(map[*ClientHandler]bool),
		Join:      make(chan *ClientHandler),
		Leave:     make(chan *ClientHandler),
		Broadcast: make(chan Message),
	}
}

func (room *Room) Run() {
	for {
		select {
		case client := <-room.Join:
			room.Clients[client] = true
			fmt.Println(client.User.Username, " joined the room ", room.Name)
		case client := <-room.Leave:
			delete(room.Clients, client)
			fmt.Println(client.User.Username, " left the room ", room.Name)
		case msg := <-room.Broadcast:
			for client := range room.Clients {
				client.Send <- msg
			}
		}
	}
}
