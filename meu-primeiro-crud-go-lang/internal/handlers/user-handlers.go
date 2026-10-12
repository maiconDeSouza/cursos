package handlers

import (
	"encoding/json"
	"meu-primeiro-crud-go-lang/internal/config"
	errapp "meu-primeiro-crud-go-lang/internal/config/err-app"
	"meu-primeiro-crud-go-lang/internal/models"
	"meu-primeiro-crud-go-lang/internal/services"
	"net/http"
	"uuid"
)

type HandlersInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
	h.cfg.Log.Info("Usuário criado com sucesso", "user", userResponse)
}

func (h *Handlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userLogin := models.UserLogin{}
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		errApp := errapp.NewBadRequestError("Json errado", err)
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao fazer login", "error", errApp)
		return
	}

	userResponse, cookie, errApp := h.services.LoginUser(userLogin)
	if errApp != nil {
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao fazer login", "error", errApp)
		return
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
	h.cfg.Log.Info("Usuário criado com sucesso", "user", userResponse)
}

func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := r.Context().Value("user_sub")
	id, err := uuid.Parse(idString.(string))
	if err != nil {
		errApp := errapp.NewBadRequestError("Uuid no formato errado", err)
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao fazer login", "error", errApp)
		return
	}

	userUpdate := models.UserUpdate{}
	if err := json.NewDecoder(r.Body).Decode(&userUpdate); err != nil {
		errApp := errapp.NewBadRequestError("Json errado", err)
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao fazer login", "error", errApp)
		return
	}

	userResponse, errApp := h.services.UpdateUser(id, userUpdate)
	if errApp != nil {
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao Atualizar os dados", "error", errApp)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
}

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := r.Context().Value("user_sub")
	id, err := uuid.Parse(idString.(string))
	if err != nil {
		errApp := errapp.NewBadRequestError("Uuid no formato errado", err)
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao fazer login", "error", errApp)
		return
	}

	if errApp := h.services.DeleteUser(id); errApp != nil {
		w.WriteHeader(errApp.Code)
		json.NewEncoder(w).Encode(errApp)
		h.cfg.Log.Error("Erro ao Atualizar os dados", "error", errApp)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}
