package services

import (
	"02-task-manager/internal/db"
	scanstring "02-task-manager/internal/scan-string"
	"02-task-manager/internal/task/models"
	"fmt"
	"strings"
)

type Services struct {
	repo db.DB
}

func (s Services) CreateTask() []models.Task {
	fmt.Println("Digite uma ou mais taks sseparadas com figurla:")
	tasks := scanstring.ScanString()
	var list []string
	var listNewTask []models.Task

	for t := range strings.SplitSeq(tasks, ",") {
		list = append(list, strings.TrimSpace(t))
	}

	for _, task := range list {
		s.repo.Index++
		newTask := models.Task{
			ID:    uint(s.repo.Index),
			Title: task,
			Done:  false,
		}
		s.repo.Tasks[newTask.ID] = newTask
		listNewTask = append(listNewTask, newTask)
	}

	return listNewTask
}

func NewServices(db db.DB) *Services {
	return &Services{
		repo: db,
	}
}
