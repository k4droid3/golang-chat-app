package main

import (
	"fmt"

	"github.com/k4droid3/golang-chat/internal/server"
)

func main() {
	sv := server.NewServer("127.0.0.1:8080", 7)
	err := sv.Start()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
