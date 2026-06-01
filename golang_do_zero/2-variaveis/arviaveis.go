package main

import "fmt"

func main() {
	var variavel1 string = "Texto qualquer"
	variavel2 := "Texto qualquer dois"

	fmt.Println(variavel1, variavel2)

	var (
		v3 string = "v3"
		v4 string = "v4"
	)
	fmt.Println(v3, v4)

	v5, v6 := "v5", "v6"

	fmt.Println(v5, v6)

	const v7 string = "v7"
	fmt.Println(v7)
}
