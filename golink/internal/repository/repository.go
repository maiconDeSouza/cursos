package repository

import "golink/internal/model"

type UserRepositoryInterface interface {
	Create(user model.User) error
	FindByEmail(email string, user *model.User) error
}
