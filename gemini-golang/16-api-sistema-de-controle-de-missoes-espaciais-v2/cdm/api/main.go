package main

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /api/v2/ships", handlers.GetAllShips)
	fmt.Println("Servidor rodando!")
	http.ListenAndServe(":2005", nil)
}
