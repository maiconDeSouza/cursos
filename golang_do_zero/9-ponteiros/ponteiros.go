package main

import "fmt"

func somaUm(n *int8) {
	*n++
}

func main() {
	var var1 int8 = 10
	var var2 int8 = var1
	var var3 *int8 = &var1

	var1++

	fmt.Println("var1 ->", var1)
	fmt.Println("var2 ->", var2)
	fmt.Println("var3 ->", var3)
	fmt.Println("var3 ->", *var3)

	somaUm(&var1)

	fmt.Println("var1 ->", var1)

}
