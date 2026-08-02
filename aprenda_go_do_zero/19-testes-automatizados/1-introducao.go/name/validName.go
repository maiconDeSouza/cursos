package name

import "strings"

func ValidName(name string) bool {
	name = strings.TrimSpace(name)
	if len(name) > 3 {
		return true
	}

	return false
}
