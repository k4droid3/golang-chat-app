package term

import (
	"os"
)

type InputHandler struct {
	Source  *os.File
	BufText chan byte
}
