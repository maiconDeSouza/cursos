package repository

import (
	"golink/internal/model"
	"uuid"
)

type UserRepositoryInterface interface {
	Create(user model.User) error
	FindByEmail(email string, user *model.User) error
}

type LinkRepositoryInterface interface {
	Create(link model.Link) error
	FindByTitle(title string, link *model.Link) error
	CountByUserID(user_id uuid.UUID) (int64, error)
}
