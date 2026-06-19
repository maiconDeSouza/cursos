package handlers

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/services"
	"encoding/json"
	"net/http"
)

func CreateShip(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	sh, err := services.CreateShip(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sh)
}

func GetAllShips(w http.ResponseWriter, r *http.Request) {
	ships := services.GetAllShips()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ships)
}

func GetAllActive(w http.ResponseWriter, r *http.Request) {
	actives := services.GetAllActives()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(actives)
}

func EditShip(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("nameShip")
	body := r.Body

	sh, err := services.EditShip(name, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sh)
}

func DeleteShip(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("nameShip")

	sh, err := services.DeleteShip(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sh)
}
