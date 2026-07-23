package main

import (
	"fmt"
	"sync"
	"time"
)

func escrevendo(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		escrevendo("Olá Mundo!")
	})

	wg.Go(func() {
		escrevendo("Programando em Go!")
	})

	wg.Wait()

}
