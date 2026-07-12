package main

import "fmt"

func main() {
	var v1 string = "V1"
	var v2 = "V2"
	v3 := "V3"

	var (
		v4 string
		v5 int
	)

	const (
		C1 = iota
		C2
		C3
	)

	fmt.Println(v1, v2, v3, v4, v5)
	fmt.Println(C1, C2, C3)
}
