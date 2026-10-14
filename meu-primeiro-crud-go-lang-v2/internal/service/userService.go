package service

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/repository"
)

type UserService struct {
	repo   repository.RepoInterface
	env    *config.Env
	errapp *errapp.ErrApp
}

func NewUserService(repo repository.RepoInterface, env *config.Env, errapp *errapp.ErrApp) *UserService {
	return &UserService{repo: repo, env: env, errapp: errapp}
}
