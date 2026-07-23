package main

import "fmt"

func main() {
	var ch = make(chan string, 2)
	ch <- "Olá Mundo!"
	ch <- "Programando em Go!"
	msg := <-ch
	msg2 := <-ch

	fmt.Println(msg)
	fmt.Println(msg2)

}
