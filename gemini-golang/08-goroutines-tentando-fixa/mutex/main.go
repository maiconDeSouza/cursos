package main

import (
	"fmt"
	"sync"
	"time"
)

var tranca sync.RWMutex
var wg sync.WaitGroup

type Cliente struct {
	name string
}

type Conta struct {
	number   string
	balance  float64
	clientes []*Cliente
}

func (c *Conta) Deposit(value float64, cliente Cliente) {
	tranca.Lock()
	if value <= 0 {
		fmt.Printf("%s você não pode depositar 0 ou menos reais!\n", cliente.name)
		tranca.Unlock()
		return
	}
	fmt.Printf("%s está fazendo um deposito de R$%.2f\n", cliente.name, value)
	fmt.Printf("aguarde...\n")
	time.Sleep(time.Duration(4) * time.Second)

	c.balance += value
	tranca.Unlock()
}

func (c *Conta) Balance() {
	fmt.Printf("Seu saldo atual é R$%.2f", c.balance)
}

func main() {
	p1 := Cliente{name: "Maria"}
	p2 := Cliente{name: "João"}

	c := Conta{number: "123", balance: 100.00, clientes: []*Cliente{&p1, &p2}}

	wg.Go(func() {
		c.Deposit(-50.00, *c.clientes[0])
	})

	wg.Go(func() {
		c.Deposit(50.00, *c.clientes[1])
	})

	wg.Wait()

	c.Balance()
}
