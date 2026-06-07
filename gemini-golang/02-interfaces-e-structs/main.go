package main

import "fmt"

type CalcularFrete interface {
	Calcular(distancia float64) float64
}

func ExibirFrete(c CalcularFrete, distancia float64) float64 {
	return c.Calcular(distancia)
}

type Correios struct{}
type Transportadora struct{}

func (c Correios) Calcular(distancia float64) float64 {
	valor := 1.5
	return distancia * valor
}

func (t Transportadora) Calcular(distancia float64) float64 {
	taxaFixa := 50.00
	valor := 0.8

	return (distancia * valor) + taxaFixa
}

func main() {
	c := Correios{}
	t := Transportadora{}
	km := 100.0

	var opcoes = map[string]float64{
		"correios":      ExibirFrete(c, km),
		"trasportadora": ExibirFrete(t, km),
	}

	fmt.Println("Opções de Frete")
	for key, value := range opcoes {
		fmt.Printf("%s: R$%.2f\n", key, value)
	}

}
