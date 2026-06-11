package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := escrever("Fala ai!!!!")

	for {
		fmt.Println(<-ch)
	}
}

func escrever(text string) <-chan string {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Go(func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Duration(2) * time.Second)
		}
	})

	return ch

}
