package main

import (
	"errors"
	"fmt"
)

func somar(a, b int) (result int) {
	result = a + b
	return
}

var f = func() {
	fmt.Println("Sou uma função chamada dentro de uma variável")
}

func ofLegalAge(age uint8) (string, error) {
	if age < 18 {
		return "", errors.New("Você não é de maior")
	}

	res := "Pode entrar, você é de maior"

	return res, nil
}

func main() {
	r := somar(23, 25)
	fmt.Println(r)
	f()

	res, err := ofLegalAge(17)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
