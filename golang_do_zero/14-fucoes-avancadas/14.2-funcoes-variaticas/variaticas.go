package main

import "fmt"

func somar(n ...int16) (total int16) {
	for _, numbers := range n {
		total += numbers
	}

	return
}

func main() {
	total := somar(23, 54, 1, 56, 23, 32, 89, 0, 31)
	fmt.Println(total)
}
