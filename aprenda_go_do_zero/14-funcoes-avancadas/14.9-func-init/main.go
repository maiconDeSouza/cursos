package main

import (
	"fmt"
	"math/rand"
	"time"
)

var i int

func generateNumber(n int) int {
	return rand.Intn(n) + 1
}

func init() {
	fmt.Println("Oi sou init, rodei primeiro!")
	time.Sleep(time.Duration(generateNumber(20)) * time.Second)
	i = generateNumber(100)
}

func main() {
	fmt.Println("Função Main!!!")
	fmt.Println(i)
}
