package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Sensor(text string, central chan string) {
	central <- text
}

func temperature(min, max int) int {
	return rand.IntN(max-min+1) + min
}

func main() {
	var central = make(chan string)

	wg.Go(func() {
		name := "Nova York"
		for {
			time.Sleep(time.Duration(24) * time.Second)
			temp := temperature(18, 26)
			text := fmt.Sprintf("hoje em %s está %d°", name, temp)
			Sensor(text, central)
		}
	})

	wg.Go(func() {
		name := "Toronto"
		for {
			time.Sleep(time.Duration(24) * time.Second)
			temp := temperature(13, 24)
			text := fmt.Sprintf("hoje em %s está %d°", name, temp)
			Sensor(text, central)
		}
	})

	wg.Go(func() {
		name := "Cidade do Mexico"
		for {
			time.Sleep(time.Duration(24) * time.Second)
			temp := temperature(14, 24)
			text := fmt.Sprintf("hoje em %s está %d°", name, temp)
			Sensor(text, central)
		}
	})

	for c := range central {
		fmt.Println(c)
	}

	wg.Wait()
}
