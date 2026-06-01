package main

import (
	"errors"
	"fmt"
)

func main() {
	var age int8 = 38
	age2 := 38 // ele vai assumir apenas int conforme seu sistema operacional 32 ou 64 bits

	fmt.Println(age, age2)

	var average float32 = 8.75

	fmt.Println(average)

	var str string = "Texto"
	char := 'A' //65 o número que ele representa na tabela ascii

	fmt.Println(str, char)

	var isvalid bool = true

	fmt.Println(isvalid)

	var err error = errors.New("Erro interno")
	fmt.Println(err)
}
