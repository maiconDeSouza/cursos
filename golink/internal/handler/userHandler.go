package handler

import (
	"encoding/json"
	"golink/internal/config"
	errapp "golink/internal/errApp"
	"golink/internal/model"
	"golink/internal/service"
	"net/http"
)

type UserHandler struct {
	service service.UserServiceInterface
	env     *config.ENV
}

func NewUserHandler(service service.UserServiceInterface, env *config.ENV) *UserHandler {
	return &UserHandler{service: service, env: env}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userRequest := model.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	userResponse, errRes := h.service.Register(&userRequest)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	login := model.Login{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		errRes := errapp.NewBadRequestError("Erro no json: "+err.Error(), err)
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	cookie, loginResponse, errRes := h.service.Login(&login)
	if errRes != nil {
		w.WriteHeader(errRes.Code)
		json.NewEncoder(w).Encode(errRes)
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loginResponse)
}
