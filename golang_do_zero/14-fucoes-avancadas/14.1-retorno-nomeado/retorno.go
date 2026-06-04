package main

import "fmt"

func calculosMatematicos(n1, n2 int8) (soma int8, subtracao int8) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, sub := calculosMatematicos(7, 23)

	fmt.Println("soma", soma)
	fmt.Println("subtração", sub)

}
