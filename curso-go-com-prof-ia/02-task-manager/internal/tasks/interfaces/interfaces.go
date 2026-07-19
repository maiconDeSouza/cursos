package interfaces

import "02-task-manager/internal/tasks/models"

type DB interface {
	IncrementID()
	GetAllTaks() []models.Task
	GetTask(id int) (*models.Task, error)
}
