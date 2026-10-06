package controller

import (
	"encoding/json"
	resterr "meu-primeiro-crud-go-lang/src/configuration/rest-err"
	"net/http"
)

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	newErr := resterr.NewBadRequestError("Requisição errada")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(newErr.Code)
	json.NewEncoder(w).Encode(newErr)
}

func FindUserByEmail(w http.ResponseWriter, r *http.Request) {}
