package handler

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/service"
)

type UserHandler struct {
	service service.UserServiceInterface
	env     *config.Env
	errapp  *errapp.ErrApp
}

func NewUserHandler(service service.UserServiceInterface, env *config.Env, errapp *errapp.ErrApp) *UserHandler {
	return &UserHandler{service: service, env: env, errapp: errapp}
}
