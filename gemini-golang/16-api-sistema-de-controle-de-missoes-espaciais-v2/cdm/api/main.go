package main

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/handlers"
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v2/ships", handlers.CreateShip)
	mux.HandleFunc("GET /api/v2/ships", handlers.GetAllShips)
	mux.HandleFunc("GET /api/v2/ships/active", handlers.GetAllActive)
	mux.HandleFunc("PUT /api/v2/ships/{nameShip}", handlers.EditShip)
	mux.HandleFunc("DELETE /api/v2/ships/{nameShip}", handlers.DeleteShip)

	fmt.Println("Servidor rodando!")
	http.ListenAndServe(":2005", middleware.Logger(mux))
}
