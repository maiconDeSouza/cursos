package repository

import "meu-primeiro-crud-go-lang-v2/internal/model"

type RepoInterface interface {
	Create(user *model.UserApp) error
}
