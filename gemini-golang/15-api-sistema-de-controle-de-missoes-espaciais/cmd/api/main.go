package main

import (
	"api-sistema-de-controle-de-missoes-espaciais/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/cmd/api/", handlers.RequestMethods)
	fmt.Println("Servidor rodando!!!")
	http.ListenAndServe(":2005", nil)
}
