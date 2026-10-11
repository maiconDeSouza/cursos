package handlers

import (
	"encoding/json"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"meu-primeiro-crud-go-lang/internal/models"
	"meu-primeiro-crud-go-lang/internal/services"
	"net/http"
)

type HandlersInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	services services.ServicesInterface
	cfg      *config.Config
}

func NewHandlers(services services.ServicesInterface, cfg *config.Config) *Handlers {
	return &Handlers{services: services, cfg: cfg}
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userRequest := models.UserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		errApp := errapp.NewBadRequestError("Json errado", err)
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao criar Usuário", "error", errApp)
		return
	}

	userResponse, errApp := h.services.CreateUser(&userRequest)
	if errApp != nil {
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao criar Usuário", "error", errApp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
	h.cfg.Log.Info("Usuário criado com sucesso", "user", userResponse)
}
