package term

import (
	"os"
)

type TermHandler struct {
	Input *inputHandler
}

type inputHandler struct {
	Source  *os.File
	BufText chan byte
}
