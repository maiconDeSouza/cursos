package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sensores = []string{"temperatura", "fumaça", "movimento"}
var sensoresMap = map[string]string{
	"temperatura": "Alerta para temperadura elevada no ambiente!",
	"fumaça":      "Alerta para fumaça no ambiente!",
	"movimento":   "Alerta para movimentação no ambiente!",
}

func monitoramento() string {

	fmt.Println("Monitorando...")
	t := rand.Intn(100) + 1
	time.Sleep(time.Duration(t) * time.Second)
	s := rand.Intn(len(sensores))

	return sensores[s]

}

func main() {
	var wg sync.WaitGroup
	chTemp, chFum, chMov := make(chan string), make(chan string), make(chan string)

	wg.Go(func() {
		for {
			s := monitoramento()
			switch s {
			case "temperatura":
				valor, ok := sensoresMap[s]
				if ok != true {
					fmt.Println("Valor não encontrada")
					continue
				}
				chTemp <- valor
			case "fumaça":
				valor, ok := sensoresMap[s]
				if ok != true {
					fmt.Println("Valor não encontrada")
					continue
				}
				chFum <- valor
			case "movimento":
				valor, ok := sensoresMap[s]
				if ok != true {
					fmt.Println("Valor não encontrada")
					continue
				}
				chMov <- valor
			default:
				continue
			}
			break
		}
	})

	select {
	case msgTemp := <-chTemp:
		fmt.Println(msgTemp)
	case msgFum := <-chFum:
		fmt.Println(msgFum)
	case msgMov := <-chMov:
		fmt.Println(msgMov)
	}

	wg.Wait()

}
