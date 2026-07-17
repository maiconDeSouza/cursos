package main

import (
	"fmt"
	"math/rand"
)

func randomNumber(n int) int {
	r := rand.Intn(n) + 1
	return r
}

func week() (r string) {
	w := randomNumber(7)
	switch w {
	case 1:
		r = "Domingo"
	case 2:
		r = "Segunda-Feira"
	case 3:
		r = "Terça-Feira"
	case 4:
		r = "Quarta-Feira"
	case 5:
		r = "Quinta-Feira"
	case 6:
		r = "Sexta-Feira"
	case 7:
		r = "Sabado"
	default:
		r = "Valor inválido"
	}
	return
}

func period() (r string) {
	n := randomNumber(3)

	switch {
	case n == 1:
		r = "Manhã"
	case n == 2:
		r = "Tarde"
	case n == 3:
		r = "Noite"
	default:
		r = "Valor inválido"
	}

	return
}

func work(w string) string {
	var r string
	switch w {
	case "Domingo", "Sabado":
		fallthrough
	case "Segunda-Feira":
		r = "Trabalho na Segunda-Feira"
	case "Terça-Feira":
		r = "Trabalho Terça-Feira"
	case "Quarta-Feira":
		r = "Trabalho Quarta-Feira"
	case "Quinta-Feira":
		r = "Trabalho Quinta-Feira"
	case "Sexta-Feira":
		r = "Trabalho Sexta-Feira"
	default:
		r = "Dia inválido"
	}
	return r
}

func main() {
	w := week()
	p := period()
	wk := work(w)

	fmt.Printf("Hoje é %s estamos no periodo %s e %s", w, p, wk)
}
