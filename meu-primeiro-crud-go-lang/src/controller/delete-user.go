package controller

import (
	"encoding/json"
	"meu-primeiro-crud-go-lang/src/model/request"
	"net/http"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := request.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
