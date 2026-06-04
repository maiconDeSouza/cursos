package main

import "fmt"

func fibonacci(p uint) uint {
	if p <= 1 {
		return p
	}
	return fibonacci(p-2) + fibonacci(p-1)
}

func main() {
	p := uint(23)

	for i := uint(0); i <= p; i++ {
		fmt.Println(fibonacci(i))
	}
}
