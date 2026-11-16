package main

import (
	"html/template"
	"os"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

var tmpl = ``

func main() {
	p := Product{
		Name:      "Caneta",
		Price:     1.23,
		Available: false,
	}

	t, err := template.New("Produto").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	if err := t.Execute(os.Stdout, p); err != nil {
		panic(err)
	}
}
