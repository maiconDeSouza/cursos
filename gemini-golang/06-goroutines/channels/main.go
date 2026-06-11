package main

import (
	"fmt"
	"sync"
	"time"
)

func CozinharMacarrao(macarrao chan<- string) {
	time.Sleep(2 * time.Second)
	macarrao <- "Macarrão"
}

func FazerMolho(molho chan<- string) {
	time.Sleep(1 * time.Second)
	molho <- "Molho"
}

func main() {
	var wg sync.WaitGroup
	macarrao, molho := make(chan string, 1), make(chan string, 1)

	wg.Go(func() {
		CozinharMacarrao(macarrao)
	})

	wg.Go(func() {
		FazerMolho(molho)
	})
	var ingrediente1 string
	var ingrediente2 string

	for i := 0; i < 2; i++ {
		select {
		case ingrediente1 = <-macarrao:
		case ingrediente2 = <-molho:
		}

	}

	wg.Wait()

	fmt.Printf("Prato %s ao %s está pronto", ingrediente1, ingrediente2)
}
