package server

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/k4droid3/golang-chat/internal/chat"
)

type Server struct {
	Addr  string
	Rooms map[string]*chat.Room
}

func NewServer(addr string, seed int64) *Server {
	return &Server{
		Addr:  addr,
		Rooms: make(map[string]*chat.Room),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	fmt.Println("Server started on", s.Addr)

	// Creating Rooms
	s.Rooms["Room0"] = chat.NewRoom("Room0")
	go s.Rooms["Room0"].Run()
	fmt.Println("Chat Room (Room0) created.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
		}

		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	botNum := r.Intn(10000)
	fmt.Printf("ClientHandler and User created for client-%d\n", botNum)
	handler := &chat.ClientHandler{
		User:       &chat.User{Username: fmt.Sprintf("client-%d", botNum), Name: fmt.Sprintf("Bot %d", botNum)},
		Connection: conn,
		Send:       make(chan chat.Message),
		Room:       s.Rooms["Room0"],
	}
	s.Rooms["Room0"].Join <- handler
	go handler.Run()
}
