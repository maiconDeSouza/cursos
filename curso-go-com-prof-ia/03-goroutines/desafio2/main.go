package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var fonte1 = []string{"dado1", "dado2", "dado3"}
var fonte2 = []string{"dado1", "dado2", "dado3", "dado4", "dado5"}
var fonte3 = []string{"dado1", "dado2"}

func gerarDados(dados []string, n int) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for _, dado := range dados {
			t := rand.Intn(10) + 1
			time.Sleep(time.Duration(t) * time.Second)
			ch <- fmt.Sprintf("fonte%d:%s", n, dado)
		}
	}()
	return ch
}

func multiplexar(ch ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	chExist := make(chan string)

	for _, c := range ch {
		wg.Go(func() {
			for v := range c {
				chExist <- v
			}
		})
	}
	go func() {
		wg.Wait()
		close(chExist)
	}()

	return chExist
}

func main() {
	f1 := gerarDados(fonte1, 1)
	f2 := gerarDados(fonte2, 2)
	f3 := gerarDados(fonte3, 3)

	chExist := multiplexar(f1, f2, f3)

	for v := range chExist {
		fmt.Println(v)
	}

}
