package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Duration(3) * time.Second)
			ch1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			ch2 <- "Canal 2"
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
