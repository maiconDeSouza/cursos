package main

import "fmt"

func fibonacci(p uint) uint {
	if p <= 1 {
		return p
	}

	return fibonacci(p-2) + fibonacci(p-1)
}

func sumDigits(n int) int {
	if n < 10 {
		return n
	}

	return (n % 10) + sumDigits(n/10)
}

func main() {
	p := uint(23)
	fmt.Println(fibonacci(p))

	n1 := 1432
	r1 := sumDigits(n1)

	n2 := 23456987
	r2 := sumDigits(n2)

	fmt.Printf("A soma dos dígitos de %d é: %d\n", n1, r1)
	fmt.Printf("A soma dos dígitos de %d é: %d\n", n2, r2)
}
