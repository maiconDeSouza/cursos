package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá goLang!")
	}()

	r := func(text string) string {
		return text
	}("Boa goLang!!!")

	f := func() {
		fmt.Println("Linguagem top")
	}

	fmt.Println(r)

	f()
}
