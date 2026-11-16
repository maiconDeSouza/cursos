package main

import (
	"html/template"
	"log"
	"os"
)

type TemplateData struct {
	Name string
}

// func main() {
// 	d := TemplateData{Name: "Dante"}

// 	t, err := template.New("test").Parse("<h1>Hello {{.Name}}</h1>")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		if err := t.Execute(os.Stdout, d); err != nil {
//			log.Fatal(err)
//		}
//	}
func main() {
	d := TemplateData{Name: "Dante"}

	t, err := template.ParseFiles("./cmd/exp/hello.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, d); err != nil {
		log.Fatal(err)
	}
}
