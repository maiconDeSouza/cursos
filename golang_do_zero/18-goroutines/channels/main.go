package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(text string, channel chan string) {
	for i := 0; i <= 5; i++ {
		channel <- text
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func main() {
	chan1 := make(chan string)
	var wg sync.WaitGroup
	init := time.Now()

	wg.Go(func() {
		go escrever("Hello, World!", chan1)
	})

	wg.Go(func() {
		escrever("Hello Goroutines!", chan1)
	})

	wg.Go(func() {
		escrever("Hello Dante!!!", chan1)
	})

	go func() {
		wg.Wait()    // Fica travada aqui até as 3 chamadas acima terminarem
		close(chan1) // Quando as 3 terminam, fecha o canal com segurança
	}()

	for m := range chan1 {
		fmt.Println(m)
	}
	wg.Wait()
	fmt.Println(time.Since(init))
}
