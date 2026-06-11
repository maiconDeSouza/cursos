package main

import "sync"

type contaCorrente struct {
	nome  string
	saldo float64
	senha string
	porta sync.Mutex
}

func (c *contaCorrente) deposito(valor float64) {
	c.porta.Lock()
	c.saldo += valor
	c.porta.Unlock()
}

func main() {
	var wg sync.WaitGroup
	c := contaCorrente{nome: "Dante", saldo: 100, senha: "123"}

	for i := 0; i <= 5; i++ {
		wg.Go(func() {
			c.deposito(50)
		})
	}

	wg.Wait()

	println("Saldo em conta:", c.saldo)

}
