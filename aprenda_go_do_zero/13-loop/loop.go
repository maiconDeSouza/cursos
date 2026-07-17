package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber(n int) int {
	r := rand.Intn(n) + 1
	return r
}

func sleepDurationSecond() {
	time.Sleep(time.Second)
}

func main() {
	i := 1

	for i <= 10 {
		sleepDurationSecond()
		fmt.Println("O Valor de i", i)
		i++
	}

	fmt.Println("-----------")
	for j := 0; j <= 10; j++ {
		sleepDurationSecond()
		fmt.Println("O Valor de j", j)
	}

	fmt.Println("-----------")
	t := 0
	r := randomNumber(10)
	for {
		sleepDurationSecond()
		t++
		if t == r {
			fmt.Println("Agora t é igual a r", t)
			break
		}
		fmt.Println("O Valor de t", t)
	}

	fmt.Println("-----------")
	for n := range 10 {
		sleepDurationSecond()
		v := n + 1
		if v == 1 {
			fmt.Println("Eu quero imprimir essa menssagem", v, "vez")
			continue
		}
		fmt.Println("Eu quero imprimir essa menssagem", v, "vezes")
	}

	fmt.Println("-----------")
	names := []string{"Dante", "Dona Maia", "Delinha"}

	for i, v := range names {
		sleepDurationSecond()
		fmt.Println(i, v)
	}

	fmt.Println("-----------")
	user := map[string]string{
		"name":     "Dante",
		"lastname": "Parrudo Kiko III",
	}

	for k, v := range user {
		sleepDurationSecond()
		fmt.Println(k, v)
	}

	for _, v := range "String" {
		sleepDurationSecond()
		fmt.Printf("%s -> tabela asceii %d\n", string(v), v)
	}
}
