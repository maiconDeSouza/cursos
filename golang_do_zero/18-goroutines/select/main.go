package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch1, ch2 := make(chan string), make(chan string)

	wg.Go(func() {
		for {
			ch1 <- "Eu sou o CH1"
			time.Sleep(time.Duration(1) * time.Second)
		}
	})

	wg.Go(func() {
		for {
			ch2 <- "Eu sou o CH2"
			time.Sleep(time.Duration(4) * time.Second)
		}
	})

	for {

		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}

	}

}
