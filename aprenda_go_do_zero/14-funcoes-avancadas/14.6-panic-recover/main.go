package main

import "fmt"

func recoverMedia() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada!", r)
	}
}

func media(a, b int) string {
	defer recoverMedia()
	m := (a + b) / 2

	if a < 0 || a > 10 {
		panic("A primeira nota tem que estar entre 0 e 10")
	}

	if b < 0 || b > 10 {
		panic("A segunda nota tem que estar entre 0 e 10")
	}

	return fmt.Sprintf("A média do aulo é %d", m)
}

func main() {
	fmt.Println(media(8, 7))
	fmt.Println(media(-1, 10))
}
