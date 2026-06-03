package main

import "fmt"

func main() {
	number := -10

	if number >= 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if newNumber := number; newNumber >= 0 {
		fmt.Println(newNumber, "-> é maior que 0")
	} else {
		fmt.Println(newNumber, "-> é menor que 0")
	}
}
