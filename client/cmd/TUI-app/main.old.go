// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"syscall"
// 	"time"
// 	"unsafe"

// 	"github.com/k4droid3/TUI-chat/internal/chat"
// )

// // func main() {
// // 	ui := tui.NewTui()
// // 	ui.Run()
// // }

// // UI Elements
// type UI struct {
// 	username  string
// 	status    string
// 	history   []chat.Message
// 	inputLine string
// 	sentLine  []string
// }

// // Render the UI (simplified)
// func (ui *UI) Render() {
// 	height := 24
// 	width := 80
// 	// Clear screen (ANSI escape code)
// 	fmt.Print("\033[2J")

// 	// Seperator
// 	fmt.Printf("\033[1;1H+")
// 	fmt.Printf("\033[1;2H%s", strings.Repeat("-", width-2))
// 	fmt.Printf("\033[1;%dH+\n", width)

// 	// Status bar
// 	fmt.Printf("\033[2;1H|")
// 	fmt.Printf("\033[2;2H%s", ui.username)
// 	fmt.Printf("\033[32m\033[2;%dH%s\033[0m", width-len(ui.status), ui.status)
// 	fmt.Printf("\033[2;%dH|\n", width)

// 	// Seperator
// 	fmt.Printf("\033[3;1H+")
// 	fmt.Printf("\033[3;2H%s", strings.Repeat("-", width-2))
// 	fmt.Printf("\033[3;%dH+\n", width)

// 	// Draw chat history (truncate to 20 lines)
// 	for i := 4; i < height-1; i++ {
// 		fmt.Printf("\033[%d;1H|", i)
// 		if i-4 < len(ui.history) {
// 			messageBubble := fmt.Sprintf("%s: %s", ui.history[i-4].User, ui.history[i-4].Content)
// 			formatedTime := ui.history[i-4].Timestamp.Format("2006-01-02 15:04:05")
// 			fmt.Printf("\033[%d;2H%s", i, messageBubble)
// 			fmt.Printf("\033[%d;%dH%s", i, width-len(formatedTime), formatedTime)
// 		}
// 		fmt.Printf("\033[%d;%dH|\n", i, width)
// 	}

// 	// Seperator
// 	fmt.Printf("\033[%d;1H+", height-1)
// 	fmt.Printf("\033[%d;2H%s", height-1, strings.Repeat("-", width-2))
// 	fmt.Printf("\033[%d;%dH+\n", height-1, width)

// 	// Draw input line
// 	fmt.Printf("\033[%d;0H> %s", height, ui.inputLine)
// }

// // Handle input (simplified)
// func (ui *UI) HandleInput() {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, _ := reader.ReadString('\n')
// 		ui.inputLine = input
// 		ui.Render()
// 	}
// }

// func TuiAppMain() {
// 	ui := &UI{
// 		username: "User: Ciri",
// 		status:   "Status: Online",
// 		history: []chat.Message{{User: "Geralt", Content: "Hmm.... a round of Gwent?", Timestamp: time.Now().Add(-20 * time.Second)},
// 			{User: "Yennefer", Content: "hey, Ciri?", Timestamp: time.Now()}},
// 		inputLine: "",
// 		sentLine:  make([]string, 0),
// 	}

// 	fd := syscall.Stdin

// 	// fmt.Println("getting original terminal settings...")
// 	// Original Terminal Settings
// 	var originalState syscall.Termios
// 	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCGETS, uintptr(unsafe.Pointer(&originalState)))
// 	if err != 0 {
// 		fmt.Printf("Error getting termios: %v\n", syscall.Errno(err))
// 		return
// 	}

// 	// fmt.Println("getting uncooked terminal state")
// 	// Uncooked
// 	newState := originalState
// 	newState.Lflag &^= syscall.ECHO | syscall.ICANON
// 	newState.Cc[syscall.VMIN] = 1
// 	newState.Cc[syscall.VTIME] = 0

// 	_, _, err = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&newState)))
// 	if err != 0 {
// 		fmt.Printf("Error setting termios: %v\n", syscall.Errno(err))
// 		return
// 	}

// 	// Reset Original Terminal
// 	defer func() {
// 		fmt.Println("\nfinal message:", ui.sentLine)
// 		if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&originalState))); err != 0 {
// 			fmt.Println("Error restoring terminal:", err)
// 		}
// 	}()

// 	// fmt.Println("now reading...")
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	reader := bufio.NewReader(os.Stdin)
// 	ui.Render()
// 	for len(c) == 0 {
// 		r, _, err := reader.ReadRune()
// 		if err != nil {
// 			break
// 		}
// 		if r == '\n' {
// 			ui.sentLine = append(ui.sentLine, ui.inputLine)
// 			if ui.inputLine == "/exit" {
// 				break
// 			}
// 			ui.inputLine = ""
// 		} else if r == '\b' || r == 127 {
// 			if len(ui.inputLine) > 0 {
// 				ui.inputLine = ui.inputLine[:len(ui.inputLine)-1]
// 			}
// 		} else {
// 			ui.inputLine += string(r)
// 		}
// 		ui.Render()
// 		time.Sleep(10 * time.Millisecond)
// 	}
// 	// if len(c) > 0 {
// 	// 	<-c
// 	// 	fmt.Println("Received interrupt, restoring terminal...")
// 	// 	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&originalState))); err != 0 {
// 	// 		fmt.Println("Error restoring terminal:", err)
// 	// 	}
// 	// }
// 	// for {
// 	// 	ui.Render()
// 	// 	time.Sleep(10 * time.Millisecond)
// 	// }
// }
