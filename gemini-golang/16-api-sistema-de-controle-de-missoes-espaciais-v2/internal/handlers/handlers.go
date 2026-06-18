package handlers

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/models"
	"encoding/json"
	"net/http"
)

func CreateShip(w http.ResponseWriter, r *http.Request) {
	sh := models.Ship{}
	body := r.Body

	err := sh.CreateShip(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sh)
}

func GetAllShips(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"Valor": "Maicon",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
