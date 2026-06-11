package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pagamento struct {
	ID    int
	Valor float64
}

func gerarPagamentos() []Pagamento {
	var pagamentos []Pagamento
	for i := 1; i < 50; i++ {
		valorAleatorio := float64(rand.Intn(990)) + 10.00
		pagamentos = append(pagamentos, Pagamento{ID: i, Valor: valorAleatorio})
	}
	return pagamentos
}

func Caixa(id int, esteira <-chan Pagamento) {
	for pagamento := range esteira {
		fmt.Printf("Caixa %d processando pagamento %d (R$ %.2f)\n", id, pagamento.ID, pagamento.Valor)
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Printf("--- Caixa %d encerrou o turno! ---\n", id)
}

func main() {
	var wg sync.WaitGroup
	esteira := make(chan Pagamento, 60)
	pagamentos := gerarPagamentos()

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			Caixa(i, esteira)
		})
	}

	for _, p := range pagamentos {
		esteira <- p
	}

	close(esteira)
	wg.Wait()
	fmt.Println("\nTodos os pagamentos foram processados com sucesso!")
}
