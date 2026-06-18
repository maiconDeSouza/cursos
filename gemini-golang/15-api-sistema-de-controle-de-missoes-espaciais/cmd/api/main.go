package main

import (
	"api-sistema-de-controle-de-missoes-espaciais/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/cmd/api/", handlers.RequestMethods)
	http.HandleFunc("/cmd/api/active", handlers.Active)
	fmt.Println("Servidor rodando!!!")
	http.ListenAndServe(":2005", nil)
}
