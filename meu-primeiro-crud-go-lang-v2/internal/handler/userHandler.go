package handler

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/service"
)

type UserHandler struct {
	service service.UserServiceInterface
	env     *config.Env
}

func NewUserHandler(service service.UserServiceInterface, env *config.Env) *UserHandler {
	return &UserHandler{service: service, env: env}
}
