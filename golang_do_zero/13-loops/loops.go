package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 0

	for n <= 10 {
		fmt.Println("n =", n)
		n++
		time.Sleep(time.Second)
	}

	for j := 0; j <= 10; j++ {
		fmt.Println("j =", j)
		time.Sleep(time.Second)
	}

	names := [...]string{"Dona Maia", "Dante", "Talu"}

	for i, name := range names {
		fmt.Println(i, name)
		time.Sleep(time.Second)
	}

	for i, letter := range "Dona Maia" {
		fmt.Println(i, string(letter))
		time.Sleep(time.Second)
	}

	dog := map[string]string{
		"name": "Dante",
		"age":  "5",
	}

	for key, value := range dog {
		fmt.Println(key, value)
		time.Sleep(time.Second)
	}

	for {
		numberRandom := rand.Intn(101)

		fmt.Println("numero sorteado foi", numberRandom)

		if numberRandom == 93 {
			fmt.Println("Parabéns o número 93 saiu")
			time.Sleep(time.Second)
			break
		}
		time.Sleep(time.Second)
	}
}
