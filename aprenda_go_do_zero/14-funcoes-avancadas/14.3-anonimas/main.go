package main

import (
	"fmt"
	"math/rand"
)

var f = func(text string) {
	fmt.Println(text)
}

func randomNumber() int {
	r := rand.Intn(1000) + 1
	return r
}

func receive(name string, f func() int) string {
	n := f()

	return fmt.Sprintf("Olá, %s! sua senha é %d", name, n)
}

func main() {
	func() {
		fmt.Println("Olá, GoLang!")
	}()

	f("ei!")

	r := receive("Dante", func() int {
		r := rand.Intn(1000) + 1
		return r
	})
	fmt.Println(r)
}
