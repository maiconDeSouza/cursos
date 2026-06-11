package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Voto struct {
	candidatoA   uint
	candidatoB   uint
	brancosNulos uint
	urna         sync.Mutex
}

func (v *Voto) Votar(vt uint) {
	v.urna.Lock()
	switch vt {
	case 1:
		v.candidatoA++
	case 2:
		v.candidatoB++
	default:
		v.brancosNulos++
	}
	v.urna.Unlock()
}

func (v *Voto) TotalDeVotos() string {
	total := v.candidatoA + v.candidatoB + v.brancosNulos
	totalValidos := v.candidatoA + v.candidatoB

	if total == 0 {
		return "Nenhum voto computado até o momento."
	}

	pA := (float64(v.candidatoA) / float64(total)) * 100
	pB := (float64(v.candidatoB) / float64(total)) * 100
	pN := (float64(v.brancosNulos) / float64(total)) * 100

	pAValidos := (float64(v.candidatoA) / float64(totalValidos)) * 100
	pBValidos := (float64(v.candidatoB) / float64(totalValidos)) * 100

	return fmt.Sprintf(
		"Resultado da Votação:\n"+
			"Candidato A: %d votos (%.2f%%)\n"+
			"Candidato B: %d votos (%.2f%%)\n"+
			"Brancos/Nulos: %d votos (%.2f%%)\n"+
			"Total de votos: %d\n"+
			"Candidato A: votos válidos (%.2f%%)\n"+
			"Candidato B: votos válidos (%.2f%%)\n",
		v.candidatoA, pA, v.candidatoB, pB, v.brancosNulos, pN, total, pAValidos, pBValidos,
	)
}

func Mesario(esteira <-chan uint, c *Voto) {
	for voto := range esteira {
		c.Votar(voto)
	}
}

func main() {
	const votos = 155_380_000
	const mesarios = 2637
	esteira := make(chan uint)

	var wg sync.WaitGroup
	c := Voto{candidatoA: 0, candidatoB: 0, brancosNulos: 0}

	for i := 1; i <= mesarios; i++ {
		wg.Go(func() {
			Mesario(esteira, &c)
		})
	}

	fmt.Println("Processando 155 milhões de votos...")
	for i := 1; i <= votos; i++ {
		numeroAleatorio := rand.Intn(3) + 1
		esteira <- uint(numeroAleatorio)
	}

	close(esteira)
	wg.Wait()

	fmt.Println(c.TotalDeVotos())
}
