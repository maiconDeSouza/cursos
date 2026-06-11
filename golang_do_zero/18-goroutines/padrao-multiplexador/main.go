package main

import (
	"fmt"
	"time"
)

func main() {
	// O multiplexar junta os dois canais infinitos em um só
	ch := multiplexar(escrever("Fala ai!!!!"), escrever("go"))

	// Lendo o canal unificado para sempre
	for {
		fmt.Println(<-ch)
	}
}

func multiplexar(ch1, ch2 <-chan string) <-chan string {
	chExit := make(chan string)

	// Como o select vai rodar para sempre em background,
	// usamos o 'go func()' direto. Sem WaitGroups locais.
	go func() {
		for {
			select {
			case msg := <-ch1:
				chExit <- msg
			case msg := <-ch2:
				chExit <- msg
			}
		}
	}()

	return chExit
}

func escrever(text string) <-chan string {
	ch := make(chan string)

	// O gerador roda para sempre em background.
	// O canal 'ch' é devolvido na hora e a goroutine gerencia o envio.
	go func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(2 * time.Second)
		}
	}()

	return ch
}
