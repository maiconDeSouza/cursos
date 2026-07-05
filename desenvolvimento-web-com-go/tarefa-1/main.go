package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 3000

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Seja bem vindo ao meu servidor HTTP com Go")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", HomeHandler)

	fmt.Printf("Servidor rodando na porta %d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))

}
