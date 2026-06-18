package handlers

import (
	"encoding/json"
	"net/http"
)

func GetAllShips(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"Valor": "Maicon",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}
