package banco

import (
	"errors"
	"fmt"
	"math/rand"
)

var conta = make(map[string]ContaBancaria)

type ContaBancaria interface {
	Sacar(valor float64) error
	Depositar(valor float64) error
	ObterSaldo() float64
	Senha(senha string) error
}

func OperacaoDeSaque(c ContaBancaria, valor float64) error {
	if valor <= 0 {
		return errors.New("você não pode valor 0 reais ou valor negativo")
	}
	err := c.Sacar(valor)
	if err != nil {
		return err
	}

	return nil
}

func OperacaoDeDeposito(c ContaBancaria, valor float64) error {
	if valor <= 0 {
		return errors.New("você não pode depositar 0 reais ou valor negativo")
	}

	return c.Depositar(valor)
}

func OperacaoDeSaldo(c ContaBancaria) float64 {
	return c.ObterSaldo()
}

func VerificarSenha(c ContaBancaria, senha string) error {
	return c.Senha(senha)
}

type contaCorrente struct {
	nome  string
	saldo float64
	senha string
}

func (c *contaCorrente) Sacar(valor float64) error {
	total := valor + 2.00

	if total > c.saldo {
		err := errors.New("Saldo insuficiente")
		return err
	}

	c.saldo -= total

	return nil
}

func (c *contaCorrente) Depositar(valor float64) error {
	c.saldo += valor

	return nil
}

func (c contaCorrente) ObterSaldo() float64 {
	return c.saldo
}

func (c contaCorrente) Senha(senha string) error {
	if c.senha != senha {
		err := errors.New("Senha inválida!")
		return err
	}

	return nil
}

type poupanca struct {
	nome  string
	saldo float64
	senha string
}

func (p *poupanca) Sacar(valor float64) error {
	if valor > p.saldo {
		err := errors.New("Saldo insuficiente")
		return err
	}

	p.saldo -= valor

	return nil
}

func (p *poupanca) Depositar(valor float64) error {
	p.saldo += valor

	return nil
}

func (p poupanca) ObterSaldo() float64 {
	return p.saldo
}

func (p poupanca) Senha(senha string) error {
	if p.senha != senha {
		err := errors.New("Senha inválida!")
		return err
	}

	return nil
}

func numeroDaConta() string {
	n1 := (rand.Intn(100) + 1)
	n2 := (rand.Intn(100) + 1)
	n3 := (rand.Intn(100) + 1)
	n4 := (rand.Intn(100) + 1)
	n5 := (rand.Intn(100) + 1)

	return fmt.Sprintf("%d%d%d%d%d", n1, n2, n3, n4, n5)

}

func CriarConta(nome string, tipoDaConta string, senha string) (map[string]ContaBancaria, string) {
	var numeroNovaConta string
	for {
		numeroNovaConta = numeroDaConta()
		_, existe := conta[numeroNovaConta]

		if !existe {
			break
		}
	}
	switch tipoDaConta {
	case "cc":
		conta[numeroNovaConta] = &contaCorrente{nome: nome, saldo: 0, senha: senha}
	case "cp":
		conta[numeroNovaConta] = &poupanca{nome: nome, saldo: 0, senha: senha}
	default:
		conta[numeroNovaConta] = &poupanca{nome: nome, saldo: 0, senha: senha}
	}
	return conta, numeroNovaConta
}
