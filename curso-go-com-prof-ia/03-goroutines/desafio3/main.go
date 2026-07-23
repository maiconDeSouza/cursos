package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Requisito 1: Canal com buffer 3 representando as vagas disponíveis
	vagas := make(chan struct{}, 3)

	// Preenche o canal com os 3 tokens iniciais (as 3 vagas livres)
	for i := 0; i < 3; i++ {
		vagas <- struct{}{}
	}

	// Requisito 2: Lança 10 goroutines representando os carros
	for i := 1; i <= 10; i++ {
		carroID := i // Captura a variável do loop para a goroutine

		wg.Go(func() {
			fmt.Printf("Carro %d aguardando vaga...\n", carroID)

			// Pega um token (bloqueia aqui se o buffer estiver vazio)
			<-vagas

			// Como o canal representa vagas *disponíveis*, as ocupadas são (3 - len)
			vagasOcupadas := 3 - len(vagas)
			fmt.Printf("Carro %d estacionou. Vagas ocupadas: %d\n", carroID, vagasOcupadas)

			// Simula tempo de permanência aleatório entre 1 e 3 segundos
			tempoEstacionado := rand.Intn(3) + 1
			time.Sleep(time.Duration(tempoEstacionado) * time.Second)

			// Devolve o token ao canal ao sair
			vagas <- struct{}{}
			fmt.Printf("Carro %d saiu.\n", carroID)
		})
	}

	// Espera todos os 10 carros terminarem o ciclo
	wg.Wait()

	fmt.Println("Todos os carros estacionaram.")
}
