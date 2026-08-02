package name

import (
	"strings"
)

func ProcessName(name string) string {
	name = strings.ToLower(name)
	nameList := strings.Split(name, " ")
	var processNameList []string

	for _, n := range nameList {
		processNameList = append(processNameList, strings.Title(n))
	}

	return strings.Join(processNameList, " ")
}
