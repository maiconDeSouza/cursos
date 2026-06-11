package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "oi"

	msg := <-ch
	fmt.Println(msg)

}
