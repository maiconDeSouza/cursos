package main

import "fmt"

func cal(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := cal(5, 23)

	fmt.Println(sum, sub)
}
