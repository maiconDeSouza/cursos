package handlers

import (
	"api-sistema-de-controle-de-missoes-espaciais/internal/models"
	"api-sistema-de-controle-de-missoes-espaciais/internal/storage"
	"encoding/json"
	"net/http"
)

func RequestMethods(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case http.MethodGet:
		getAllShips(w, r)
	case http.MethodPost:
		createShip(w, r)
	case http.MethodPut:
		upShip(w, r)
	case http.MethodDelete:
	default:
		http.Error(w, "Método não encontrado", http.StatusMethodNotAllowed)
	}
}

func createShip(w http.ResponseWriter, r *http.Request) {
	sh := models.Ship{}

	if err := json.NewDecoder(r.Body).Decode(&sh); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	st := storage.StoageShips
	st.AddShip(&sh)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sh)
}

func getAllShips(w http.ResponseWriter, r *http.Request) {
	st := storage.StoageShips

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(st.GetSships())
}

func upShip(w http.ResponseWriter, r *http.Request) {
	newSh := models.Ship{}
	st := storage.StoageShips

	if err := json.NewDecoder(r.Body).Decode(&newSh); err != nil {
		http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
		return
	}

	shipName := r.URL.Query().Get("name")
	sh, err := st.ShipExists(shipName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	st.UPship(newSh.Name, newSh.Status, newSh.Active, sh)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sh)
}
