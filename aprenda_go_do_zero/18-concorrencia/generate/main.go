package main

import (
	"fmt"
	"time"
)

func escrever(txt string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			time.Sleep(time.Duration(3) * time.Second)
			ch <- fmt.Sprintf("Valor recebido: %s", txt)
		}

	}()

	return ch
}

func main() {
	ch := escrever("Olá, Mundo!")

	for c := range ch {
		fmt.Println(c)
	}
}
