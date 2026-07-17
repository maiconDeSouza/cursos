package main

import (
	"fmt"
	"math/rand"
)

func randomNumber() int {
	n := rand.Intn(3) + 1
	return n
}

func main() {
	age := 23

	if age >= 60 {
		fmt.Println("Terceira idade")
	} else if age >= 18 {
		fmt.Println("Adulto")
	} else {
		fmt.Println("de menor")
	}

	q := randomNumber()

	if n := randomNumber(); n == q {
		fmt.Println("Você acertou!!! O número é", n)
	} else {
		fmt.Println("Você errou!!! O número não é", n)
	}

	user := map[string]string{
		"name":  "Dante",
		"breed": "Pug",
	}

	k := "breed"

	if v, exist := user[k]; exist {
		fmt.Println("O valor existe e é", v)
	}
}
