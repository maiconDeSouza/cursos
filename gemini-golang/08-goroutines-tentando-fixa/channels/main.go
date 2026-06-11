package main

import (
	"fmt"
	"sync"
	"time"
)

func Maria(canal chan string) {
	fmt.Println("Maria está preparando o bolo...")
	time.Sleep(2 * time.Second)
	// Como a Maria envia a string "Bolo de chocolate" para o canal aqui?
	canal <- "Bolo de chocolate"
}

func main() {
	var wg sync.WaitGroup
	// 1. Como você cria o canal de strings chamado 'esteira'?
	esteira := make(chan string)
	// 2. Chame a função Maria usando uma Goroutine (passando a esteira)
	wg.Go(func() {
		Maria(esteira)
	})
	fmt.Println("João está esperando o bolo...")

	// 3. Como o João tira a mensagem da esteira e guarda na variável 'presente'?
	presente := <-esteira

	wg.Wait()

	fmt.Println("João recebeu:", presente)
}
