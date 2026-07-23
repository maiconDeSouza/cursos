package main

import (
	"fmt"
	"sync"
	"time"
)

func escrevendo(txt string, ch chan string, duration int) {
	for i := 0; i < duration; i++ {
		ch <- txt
		time.Sleep(time.Second * time.Duration(duration))
	}
}

func main() {
	var wg sync.WaitGroup

	var ch = make(chan string)

	wg.Go(func() {
		escrevendo("Olá Mundo!", ch, 5)
	})

	wg.Go(func() {
		escrevendo("Programando em Go!", ch, 3)
	})

	go func() {
		wg.Wait()

		close(ch)
	}()

	for txt := range ch {
		fmt.Println(txt)
	}

}
