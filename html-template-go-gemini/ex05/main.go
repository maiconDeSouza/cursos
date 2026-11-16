package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const PORT = 2323

type Handler struct {
	Template *template.Template
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {

	if err := h.Template.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	t, err := template.ParseFiles("./ex05/index.html")
	if err != nil {
		panic(err)
	}

	h := Handler{Template: t}

	mux.HandleFunc("GET /", h.Home)

	fmt.Printf("🚀 Servidor iniciado com sucesso na porta: %d", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux); err != nil {
		log.Fatal(err)
	}
}
