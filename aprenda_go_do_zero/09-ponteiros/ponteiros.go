package main

import "fmt"

func main() {
	var n1 int
	var n2 int = n1
	var n3 *int

	fmt.Println("Valor de n1 e n2 e n3 ->", n1, n2, n3) // Valor de n1 e n2 e n3 -> 0 0 <nil>

	n1 = 23
	n3 = &n1

	fmt.Println("Valor de n1 e n2 e n3 ->", n1, n2, n3)  // Valor de n1 e n2 e n3 -> 23 0 0x1a2615a08008
	fmt.Println("Valor de n1 e n2 e n3 ->", n1, n2, *n3) // Valor de n1 e n2 e n3 -> 23 0 23

	*n3++
	fmt.Println("Valor de n1 e n2 e n3 ->", n1, n2, n3)  // Valor de n1 e n2 e n3 -> 24 0 0x1a2615a08008
	fmt.Println("Valor de n1 e n2 e n3 ->", n1, n2, *n3) // Valor de n1 e n2 e n3 -> 24 0 24
}
