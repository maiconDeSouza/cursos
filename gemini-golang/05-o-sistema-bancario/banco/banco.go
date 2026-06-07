package banco

import "errors"

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

type ContaCorrente struct {
	nome  string
	saldo float64
	senha string
}

func (c *ContaCorrente) Sacar(valor float64) error {
	total := valor + 2.00

	if total > c.saldo {
		err := errors.New("Saldo insuficiente")
		return err
	}

	c.saldo -= total

	return nil
}

type Poupanca struct {
	nome  string
	saldo float64
	senha string
}
