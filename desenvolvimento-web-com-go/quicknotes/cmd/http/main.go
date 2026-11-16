package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Name string
}

func path(p string) string {
	return fmt.Sprintf("./view/templates/%s", p)
}

func Home(w http.ResponseWriter, r *http.Request) {
	d := TemplateData{Name: "Dante"}

	t, err := template.ParseFiles(path("home.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := t.Execute(w, d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func View(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Lista uma tarefa")
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Criar uma tarefa")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", Home)
	mux.HandleFunc("GET /view", View)
	mux.HandleFunc("POST /create", Create)

	fmt.Printf("🚀 Servidor iniciado com sucesso na porta: %d", 2008)
	log.Fatal(http.ListenAndServe(":2008", mux))
}
