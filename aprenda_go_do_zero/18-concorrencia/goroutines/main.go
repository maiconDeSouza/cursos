package main

import (
	"fmt"
	"time"
)

func escrevendo(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	// escrevendo("Olá Mundo!")
	// escrevendo("Programando em Go!")

	go escrevendo("Olá Mundo!")
	escrevendo("Programando em Go!")
}
