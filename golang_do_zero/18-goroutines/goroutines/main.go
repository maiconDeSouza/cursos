package main

import (
	"fmt"
	"time"
)

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Duration(5) * time.Second)
	}
}

func main() {
	go escrever("Hello, World!")
	go escrever("Hello Goroutines!")
	go escrever("Hello Dante!!!")
	time.Sleep(time.Duration(5) * time.Second)

}
