package main

import "fmt"

// 1. O Contrato (Interface)
type Finalizar interface {
	end() string
}

// 2. A Função Central (Blindada)
// Ela não quer saber se é Pix ou Cartão, apenas se a estrutura possui o método end()
func finalizarCompra(f Finalizar) string {
	return f.end()
}

// 3. Meio de Pagamento 1: Pix
type Pix struct {
	Name  string
	Value float64
}

func (p Pix) end() string {
	return fmt.Sprintf("Finalizando a compra no valor de R$%.2f com %s", p.Value, p.Name)
}

// 4. Meio de Pagamento 2: Cartão (Novo!)
type Cartao struct {
	Name     string
	Value    float64
	Parcelas int
}

// O Cartão implementa o contrato end() com a sua própria regra de negócio (parcelamento)
func (c Cartao) end() string {
	valorParcela := c.Value / float64(c.Parcelas)
	return fmt.Sprintf("Finalizando a compra de R$%.2f no %s em %dx de R$%.2f", c.Value, c.Name, c.Parcelas, valorParcela)
}

func main() {
	// Simulando o Checkout no sistema
	pagamentoPix := Pix{Name: "Pix", Value: 150.00}
	pagamentoCartao := Cartao{Name: "Cartão de Crédito", Value: 300.00, Parcelas: 3}

	// A MÁGICA ACONTECE AQUI:
	// A mesma função aceita os dois structs sem dar erro de tipagem
	fmt.Println(finalizarCompra(pagamentoPix))
	fmt.Println(finalizarCompra(pagamentoCartao))
}
