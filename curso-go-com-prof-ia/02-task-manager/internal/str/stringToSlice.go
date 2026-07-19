package str

import "strings"

func StringToSlice(str string) []string {
	sliceString := strings.Split(str, ",")
	sliceClear := []string{}

	for _, str := range sliceString {
		sliceClear = append(sliceClear, strings.TrimSpace(str))
	}

	return sliceClear
}
