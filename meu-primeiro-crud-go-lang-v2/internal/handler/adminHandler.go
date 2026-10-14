package handler

import (
	"encoding/json"
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/model"
	"meu-primeiro-crud-go-lang-v2/internal/service"
	"net/http"
)

type AdminHandler struct {
	service service.AdminServiceInterface
	env     *config.Env
	errapp  *errapp.ErrApp
}

func NewAdminHandler(service service.AdminServiceInterface, env *config.Env, errapp *errapp.ErrApp) *AdminHandler {
	return &AdminHandler{service: service, env: env, errapp: errapp}
}

func (h *AdminHandler) NewAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	adminRequest := model.Request{}
	if err := json.NewDecoder(r.Body).Decode(&adminRequest); err != nil {
		errRes := h.errapp.BadRequestError("Erro no json", err)
		w.WriteHeader(errRes)
		json.NewEncoder(w).Encode()
	}
}
