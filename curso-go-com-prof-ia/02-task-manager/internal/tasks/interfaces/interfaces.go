package interfaces

import "02-task-manager/internal/tasks/models"

type DB interface {
	GetAllTaks() []models.Task
	GetTask(id int) (*models.Task, error)
	Salve(task *models.Task) models.Task
	DoneTask(id int) error
}
