package main

import "fmt"

func sum(value ...int) (s int) {
	for _, v := range value {
		s += v
	}
	return
}

func main() {
	s := sum(63, -72, -94, 89, -30, -38, -43, -65, 88, -74, 73, 89, 39, -78, 51, 8, -92, -93, -77, -45, -41, 29, 54, -94, 43, -50, 83, 66, 79, 39)
	fmt.Println(s)
}
