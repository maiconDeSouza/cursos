package routes

import (
	"meu-primeiro-crud-go-lang/src/controller"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /user/id/{id}", controller.FindUserByID)
	mux.HandleFunc("GET /user/email/{email}", controller.FindUserByEmail)
	mux.HandleFunc("POST /user", controller.CreateUser)
	mux.HandleFunc("PUT /user/id/{id}", controller.UpdateUser)
	mux.HandleFunc("DELETE /user/id/{id}", controller.DeleteUser)

}
