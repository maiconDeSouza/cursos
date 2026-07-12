package main

import (
	"errors"
	"fmt"
)

func main() {
	// existem 5 tipos INT em go -> int, int8, int16, int32, int64
	// Você declarando paenas int ele vai assumir o tamanho entre int32 e int64 de
	// acordo com a arquiterua do seu sistema operacional
	// A diferença é de acordo com quantidade de bits cada um suporta
	var n1 int = -239864
	var n2 int = 239864
	fmt.Println(n1, n2)

	//uint -> suporta apenas números positivos
	var n3 uint = 23
	fmt.Println(n3)

	//alias
	//int 32 -> rune -> usado em caracteres para a tabela ascii
	//uint 8 -> byte

	var n4 rune = 23456
	var n5 byte = 23
	fmt.Println(n4, n5)

	// números reais float32 e flat64
	var (
		n6 float32 = 35.87
		n7 float64 = 38908.56
	)
	fmt.Println(n6, n7)

	// Se você declarar um float apenas por inferencia ele sempre vai assumir float apenas
	// Mas vocẽ nunca pode declarar apenas como flaot, tem que usar float32 ou float64
	n8 := 23.23
	fmt.Println(n8)

	//tipo string é um conjunto de caracteres
	s := "texto"
	var s2 string = "texto2"
	fmt.Println(s, s2)

	// a mais proxímo que go tem com o tipo char é declarar as variáveis com aspas simples
	char := 'A'  //fica com o tipo rune
	char2 := 'a' //fica com o tipo rune
	char3 := '%' //fica com o tipo rune
	// char4 := '%a' -> não aceita mais de um caracter

	fmt.Println(char)  //65 o valor dele na tabela ascii
	fmt.Println(char2) //97 o valor dele na tabela ascii
	fmt.Println(char3) //37 o valor dele na tabela ascii

	//Tipo boleano apenas true ou falso
	b := true
	fmt.Println(b)

	var err error
	fmt.Println(err) //tipo erro em go, valor zero (inicial) é igual a nil (nulo)

	var err2 error = errors.New("Erro interno")
	fmt.Println(err2) //para atribuir um erro é necessário usar o pacote do go errors

}
