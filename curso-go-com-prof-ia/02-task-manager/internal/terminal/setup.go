package terminal

import (
	"bufio"
	"os"
)

var globalScanner = bufio.NewScanner(os.Stdin)
