package main

import (
	"fmt"

	"github.com/k4droid3/TUI-chat/internal/client"
)

func main() {
	c := client.NewClient("Ciri")
	if err := c.Start(); err != nil {
		fmt.Println("Error Starting Client.")
		panic(err)
	}
}
