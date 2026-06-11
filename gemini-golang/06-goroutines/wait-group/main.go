package main

import (
	"fmt"
	"sync"
	"time"
)

func CozinharMacarrao() {
	time.Sleep(2 * time.Second)
	fmt.Println("Macarrão pronto!")
}

func FazerMolho() {
	time.Sleep(1 * time.Second)
	fmt.Println("Molho pronto!")
}

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		CozinharMacarrao()
	})

	wg.Go(func() {
		FazerMolho()
	})

	wg.Wait()
}
