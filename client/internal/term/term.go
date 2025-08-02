package term

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

type TermHandler struct {
	InputSource     *os.File
	BufChar         chan rune
	ogState         syscall.Termios
	newState        syscall.Termios
	fd              int // unix file descriptor
	InterruptSignal chan os.Signal
}

func NewTermHandler() *TermHandler {
	return &TermHandler{
		InputSource:     os.Stdin,
		BufChar:         make(chan rune),
		ogState:         syscall.Termios{},
		newState:        syscall.Termios{},
		fd:              syscall.Stdin,
		InterruptSignal: make(chan os.Signal, 1),
	}
}

func (t *TermHandler) getOgState() error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCGETS, uintptr(unsafe.Pointer(&t.ogState))); err != 0 {
		return err
	}
	return nil
}

func (t *TermHandler) setNewState() error {
	t.newState = t.ogState
	t.newState.Lflag &^= syscall.ECHO | syscall.ICANON
	t.newState.Cc[syscall.VMIN] = 1
	t.newState.Cc[syscall.VTIME] = 0
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCSETS, uintptr(unsafe.Pointer(&t.newState))); err != 0 {
		return err
	}
	return nil
}

func (t *TermHandler) setOgState() error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(t.fd), syscall.TCSETS, uintptr(unsafe.Pointer(&t.ogState))); err != 0 {
		return err
	}
	return nil
}

func (t *TermHandler) Start() error {
	fmt.Println("Starting Terminal Handler...")
	if err := t.getOgState(); err != nil {
		fmt.Println("Failed to get Original State.")
		return err
	}
	if err := t.setNewState(); err != nil {
		fmt.Println("Failed to change terminal to raw mode.")
		t.setOgState()
		return err
	}

	go signal.Notify(t.InterruptSignal, os.Interrupt)

	go func() {
		reader := bufio.NewReader(t.InputSource)
		for {
			r, _, err := reader.ReadRune()
			if err != nil {
				fmt.Println("Failed to read rune.")
				panic(err)
			}
			t.BufChar <- r
		}
	}()

	return nil
}

func (t *TermHandler) Stop() error {
	fmt.Println("Stopping Terminal Handler...")
	if err := t.setOgState(); err != nil {
		fmt.Println("Failed to reset Original State.")
		return err
	}
	return nil
}
