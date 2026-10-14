package service

import (
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/model"
)

type UserServiceInterface interface {
	NewAdmin(adminRequest model.Request) (*model.Response, *errapp.ErrAppResult)
}

type AdminServiceInterface interface {
}
