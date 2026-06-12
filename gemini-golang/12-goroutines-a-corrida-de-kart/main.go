package main

import (
	"fmt"
	"math/rand"
	"time"
)

func temp() {
	t := 120 + rand.Intn(121)
	time.Sleep(time.Duration(t) * time.Second)
}

func Kart(name string, finishline chan string) {
	temp()
	finishline <- name
}

func main() {
	finishline := make(chan string, 4)

	go Kart("Mario", finishline)
	go Kart("Luigi", finishline)
	go Kart("Peach", finishline)
	go Kart("Bowser", finishline)

	fmt.Printf("%s recebe a bandeirada!", <-finishline)

}
