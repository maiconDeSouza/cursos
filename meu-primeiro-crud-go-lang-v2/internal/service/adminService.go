package service

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/repository"
)

type AdminService struct {
	repo repository.RepoInterface
	env  *config.Env
}

func NewAdminService(repo repository.RepoInterface, env *config.Env) *AdminService {
	return &AdminService{repo: repo, env: env}
}
