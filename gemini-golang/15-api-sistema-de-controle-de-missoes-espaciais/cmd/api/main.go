package main

import (
	"api-sistema-de-controle-de-missoes-espaciais/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/ships/", handlers.RequestMethods)
	http.HandleFunc("/api/v1/ships/active", handlers.Active)
	fmt.Println("Servidor rodando!!!")
	http.ListenAndServe(":2005", nil)
}
