package main

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/handlers"
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("POST /api/v2/ships", middleware.Logger(handlers.CreateShip))
	http.HandleFunc("GET /api/v2/ships", handlers.GetAllShips)
	http.HandleFunc("GET /api/v2/ships/active", handlers.GetAllActive)
	http.HandleFunc("PUT /api/v2/ships/{nameShip}", handlers.EditShip)
	http.HandleFunc("DELETE /api/v2/ships/{nameShip}", handlers.DeleteShip)

	fmt.Println("Servidor rodando!")
	http.ListenAndServe(":2005", nil)
}
