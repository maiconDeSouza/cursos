package main

import "fmt"

type Produto interface {
	Detalhes() string
}

type Livro struct {
	Titulo string
	Autor  string
}

func (l Livro) Detalhes() string {
	return fmt.Sprintf(
		"Livro: %s escrito por %s",
		l.Titulo,
		l.Autor,
	)
}

type Eletronico struct {
	Modelo   string
	Voltagem int
}

func (e Eletronico) Detalhes() string {
	return fmt.Sprintf(
		"Eletrônico: %s - %dv",
		e.Modelo,
		e.Voltagem,
	)
}

func main() {
	catalogo := map[string]Produto{
		"LIV01": Livro{
			Titulo: "O Senhor dos Anéis",
			Autor:  "J.R.R. Tolkien",
		},
		"ELE99": Eletronico{
			Modelo:   "PlayStation 5",
			Voltagem: 220,
		},
	}

	for codigo, produto := range catalogo {
		fmt.Printf("%s -> %s\n", codigo, produto.Detalhes())
	}
}
