package main

import (
	"fmt"
	"math/rand"
)

type Bandeira string

const (
	Visa       Bandeira = "Visa"
	Mastercard Bandeira = "Mastercard"
	Nenhuma    Bandeira = ""
)

// Interface com método que recebe o valor como parâmetro
type Pagamento interface {
	Pagar(valor float64) string
}

// Função que recebe um pagamento e o valor a ser pago
func Processar(p Pagamento, nome string, valor float64) {
	fmt.Println(nome, p.Pagar(valor))
}

type Cliente struct {
	nome             string
	valor            float64
	formaDePagamento string
	bandeira         Bandeira
}

// Pix – não armazena valor, só a implementação do método
type Pix struct{}

func (p Pix) Pagar(valor float64) string {
	valorFinal := valor * 0.98
	return fmt.Sprintf("Pagamento de R$ %.2f via Pix realizado com desconto de 2%%. Valor pago: R$ %.2f",
		valor, valorFinal)
}

// Cartao – armazena apenas a bandeira
type Cartao struct {
	bandeira Bandeira
}

func (c Cartao) Pagar(valor float64) string {
	valorFinal := valor * 1.05
	return fmt.Sprintf("Pagamento de R$ %.2f no cartão %s com acréscimo de 5%%. Valor pago: R$ %.2f",
		valor, c.bandeira, valorFinal)
}

// Boleto – armazena o código de barras (gerado uma vez)
type Boleto struct {
	codigo string
}

// Gera um código de barras fictício (exemplo)
func gerarCodigoBarras() string {
	// apenas para demonstração
	return fmt.Sprintf("%d%d%d.%d%d%d", rand.Intn(10), rand.Intn(10), rand.Intn(10),
		rand.Intn(10), rand.Intn(10), rand.Intn(10))
}

func (b Boleto) Pagar(valor float64) string {
	valorFinal := valor
	if valor > 100 { // corrigido: maior que 100
		valorFinal = valor * 0.99
	}
	return fmt.Sprintf("Pagamento de R$ %.2f via Boleto (código: %s). Valor pago: R$ %.2f",
		valor, b.codigo, valorFinal)
}

func main() {
	clientes := map[uint]Cliente{
		1:  {nome: "Alice Silva", valor: 150.50, formaDePagamento: "Cartão", bandeira: Visa},
		2:  {nome: "Bruno Costa", valor: 89.90, formaDePagamento: "Cartão", bandeira: Mastercard},
		3:  {nome: "Carla Souza", valor: 420.00, formaDePagamento: "Pix", bandeira: Nenhuma},
		4:  {nome: "Diego Oliveira", valor: 215.30, formaDePagamento: "Boleto", bandeira: Nenhuma},
		5:  {nome: "Elena Santos", valor: 95.00, formaDePagamento: "Cartão", bandeira: Visa},
		6:  {nome: "Fernando Lima", valor: 310.00, formaDePagamento: "Pix", bandeira: Nenhuma},
		7:  {nome: "Gabriela Rocha", valor: 1250.00, formaDePagamento: "Cartão", bandeira: Visa},
		8:  {nome: "Henrique Alves", valor: 65.40, formaDePagamento: "Boleto", bandeira: Nenhuma},
		9:  {nome: "Isabela Mendes", valor: 540.80, formaDePagamento: "Cartão", bandeira: Mastercard},
		10: {nome: "João Pereira", valor: 18.90, formaDePagamento: "Pix", bandeira: Nenhuma},
	}

	for _, c := range clientes {
		if c.formaDePagamento == "Cartão" {
			cartao := Cartao{bandeira: c.bandeira}
			Processar(cartao, c.nome, c.valor)
			continue
		}

		if c.formaDePagamento == "Pix" {
			pix := Pix{}
			Processar(pix, c.nome, c.valor)
			continue
		}

		if c.formaDePagamento == "Boleto" {
			boleto := Boleto{codigo: gerarCodigoBarras()}
			Processar(boleto, c.nome, c.valor)
			continue
		}
	}
}
