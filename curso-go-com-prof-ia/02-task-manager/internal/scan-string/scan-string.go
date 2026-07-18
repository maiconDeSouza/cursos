package scanstring

import (
	"bufio"
	"os"
	"strings"
)

func ScanString() string {
	sc := bufio.NewScanner(os.Stdin)

	if sc.Scan() {
		input := sc.Text()

		input = strings.TrimSpace(input)

		return input
	}
	return ""
}
