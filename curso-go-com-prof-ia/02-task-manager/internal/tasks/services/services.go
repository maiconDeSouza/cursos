package services

import (
	"02-task-manager/internal/tasks/interfaces"
	"02-task-manager/internal/tasks/repositories"
)

type Services struct {
	repo interfaces.DB
}

func (s Services) Qualquer() {
	s.repo.IncrementID()
}

func NewServices(db repositories.DB) *Services {
	return &Services{
		repo: &db,
	}
}
