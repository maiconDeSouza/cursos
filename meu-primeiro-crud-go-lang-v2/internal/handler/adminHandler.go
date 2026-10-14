package handler

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/service"
)

type AdminHandler struct {
	service service.AdminServiceInterface
	env     *config.Env
}

func NewAdminHandler(service service.AdminServiceInterface, env *config.Env) *AdminHandler {
	return &AdminHandler{service: service, env: env}
}
