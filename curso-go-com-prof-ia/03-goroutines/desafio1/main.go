package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	channelTick, channelCancel := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Duration(500) * time.Millisecond)
			channelTick <- "tick"
		}
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			channelCancel <- "Relógio parado pelo usuário"
		}
	}()

	for {
		select {
		case msg := <-channelTick:
			fmt.Println(msg)
		case msg := <-channelCancel:
			fmt.Println(msg)
			return
		}
	}

}
