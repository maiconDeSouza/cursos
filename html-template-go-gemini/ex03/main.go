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

type Store struct {
	Products []Product
}

var tmpl = `<ul>
    {{range .Products}}
        <li>{{.Name}} - R${{.Price}}</li> 
    {{else}}
        <li>Nenhum produto cadastrado</li>
    {{end}}
    </ul>`

func main() {
	products := []Product{
		{
			Name:      "Notebook Dell Inspiron",
			Price:     3500.00,
			Available: true,
		},
		{
			Name:      "Mouse Sem Fio Logitech",
			Price:     120.50,
			Available: true,
		},
		{
			Name:      "Monitor Ultrawide 29\"",
			Price:     1199.90,
			Available: false,
		},
	}

	store := Store{
		Products: products,
	}

	t, err := template.New("produtos").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	if err := t.Execute(os.Stdout, store); err != nil {
		panic(err)
	}
}
