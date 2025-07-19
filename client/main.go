package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the server.")
	defer conn.Close()

	conn.Write([]byte("Hello please let me join.\n"))

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error reading:", err)
				return
			}
			fmt.Println("Server response: ", string(buf[:n]))
		}
	}()

	for {
		var input string
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("Exiting...")
			return
		}
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
		fmt.Println("Sent to server:", input)
	}
}
