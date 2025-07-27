package main

import "fmt"

var a int32 = 21

func main() {
	var b int8 = 2
	result := a + int32(b)

	fmt.Printf("O valor de result é %v que é um %T", result, result)
}
