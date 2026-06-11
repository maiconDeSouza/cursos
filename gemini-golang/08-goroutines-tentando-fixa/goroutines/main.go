package main

import (
	"fmt"
	"time"
)

func Falar(pessoa string) {
	fmt.Println(pessoa, "começou a falar...")
	time.Sleep(time.Second)
	fmt.Println(pessoa, "terminou!")
}

func main() {
	go Falar("Maria")
	Falar("João")
}
