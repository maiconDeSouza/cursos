package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() string
}

func escrever(f forma) {
	fmt.Printf("%s\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() string {
	a := r.altura * r.largura

	return fmt.Sprintf("A área do retangulo é %.2f", a)
}

type circulo struct {
	raio float64
}

func (c circulo) area() string {
	a := math.Pi * (c.raio * c.raio)

	return fmt.Sprintf("A área do circulo é %.2f", a)
}

func main() {
	r := retangulo{altura: 2.83, largura: 5.23}
	c := circulo{raio: 3.14}

	escrever(r)
	escrever(c)
}
