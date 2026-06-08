package banco

import (
	"errors"
	"fmt"
	"math/rand"
)

var conta = make(map[string]ContaBancaria)

type senhaGerencia struct {
	senha string
}

func (s senhaGerencia) LoginGerencia(senha string) error {
	if s.senha != senha {
		err := errors.New("Senha errada!")
		return err
	}
	return nil
}

var Gerencia = senhaGerencia{senha: "1223"}

type ContaBancaria interface {
	Sacar(valor float64) error
	Depositar(valor float64) error
	ObterSaldo() float64
	Senha(senha string) error
	PegarNome() string
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

func VerificarNome(c ContaBancaria) string {
	return c.PegarNome()
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

func (c contaCorrente) PegarNome() string {
	return c.nome
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

func (p poupanca) PegarNome() string {
	return p.nome
}

func numeroDaConta() string {
	n1 := (rand.Intn(10))
	n2 := (rand.Intn(10))
	n3 := (rand.Intn(10))

	return fmt.Sprintf("%d%d%d", n1, n2, n3)

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

func BuscarConta(numero string) (ContaBancaria, error) {
	if len(numero) != 3 {
		err := errors.New("O número da conta deve ter 3 digitos!")
		return nil, err
	}

	conta, existe := conta[numero]
	if !existe {
		err := errors.New("Conta não encontrada")
		return nil, err
	}
	return conta, nil
}

func TodasContas() {
	for key, value := range conta {
		fmt.Printf("Conta: [%s] - Nome: [%s] - Saldo: [%.2f]\n", key, value.PegarNome(), value.ObterSaldo())
	}
}
