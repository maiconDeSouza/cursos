package main

import (
	"fmt"
	"sync"
	"time"
)

// Passo 1: Struct para identificar o resultado de cada posição
type Resultado struct {
	posicao uint
	valor   uint
}

func fibonacci(p uint) uint {
	if p <= 1 {
		return p
	}
	return fibonacci(p-2) + fibonacci(p-1)
}

// Passo 2: O Worker que consome a esteira de tarefas
func worker(jobs <-chan uint, results chan<- Resultado) {
	for n := range jobs {
		res := fibonacci(n)
		results <- Resultado{posicao: n, valor: res}
	}
}

func main() {
	init := time.Now()
	p := uint(50)

	// Canais com capacidade para aguentar todas as tarefas
	jobs := make(chan uint, p+1)
	results := make(chan Resultado, p+1)

	var wg sync.WaitGroup

	// Passo 3: Disparar o Pool de Workers (vamos contratar 4 trabalhadores)
	numWorkers := 4
	for w := 1; w <= numWorkers; w++ {
		wg.Go(func() {
			worker(jobs, results)
		})
	}

	// Passo 4: Alimentar a esteira de Jobs e fechar o canal em seguida
	for i := uint(0); i <= p; i++ {
		jobs <- i
	}
	close(jobs) // Avisa os workers que não há mais novas posições vindo

	// Passo 5: O Gerente que espera os workers e fecha o canal de resultados
	go func() {
		wg.Wait()
		close(results)
	}()

	// Passo 6: Consumir e exibir os resultados na tela
	for r := range results {
		fmt.Printf("Fibonacci de %d é %d\n", r.posicao, r.valor)
	}

	fmt.Printf("tempo total -> %s\n", time.Since(init))
}
