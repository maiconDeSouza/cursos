package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(text string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func main() {
	init := time.Now()
	var wg sync.WaitGroup

	// wg.Go substitui simultaneamente o wg.Add(1), a palavra 'go' e o 'defer wg.Done()'
	wg.Go(func() {
		escrever("Hello, World!")
	})

	wg.Go(func() {
		escrever("Hello Goroutines!")
	})

	wg.Go(func() {
		escrever("Hello Dante!!!")
	})

	// Continua bloqueando a execução até que todas as funções passem pelo wg.Go
	wg.Wait()
	fmt.Println(time.Since(init))
}
