package main

import "fmt"

// Operadores aritimeticos
// + soma
// - subtração
// * multiplicação
// / divisão
// % modulo

// Operadores de atribuição
// =
// :=

// Operadores de comparação
// >=
// >
// <
// <=
// ==
// !=

// Operadores lógicos
// &&
// ||

// Operadores unários
// variavel++ -> no go só da para fazer o pós
// variavel += 10

// go não tem operadores ternarios

var n1 int8 = 8
var n2 int16 = 15

func main() {
	somar := int16(n1) + n2
	fmt.Println(somar)
}
