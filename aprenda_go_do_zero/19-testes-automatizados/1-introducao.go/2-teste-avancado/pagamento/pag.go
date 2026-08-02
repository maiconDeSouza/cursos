package pagamento

type Pagmento interface {
	Taxa(valor float64) float64
}

func Taxa(p Pagmento, value float64) float64 {
	return p.Taxa(value)
}

type Cartao struct {
	taxa float64
}

func (c *Cartao) Taxa(valor float64) float64 {
	desconto := valor * (c.taxa / 100)

	return valor - desconto
}

type Boleto struct {
	taxa float64
}

func (b *Boleto) Taxa(valor float64) float64 {
	desconto := valor * (b.taxa / 100)

	return valor - desconto
}

func NewPagamento() (*Cartao, *Boleto) {
	c := &Cartao{taxa: 5}
	b := &Boleto{taxa: 1}

	return c, b
}
