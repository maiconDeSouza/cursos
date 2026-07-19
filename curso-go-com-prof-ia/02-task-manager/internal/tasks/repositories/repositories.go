package repositories

import (
	"02-task-manager/internal/tasks/models"
	"cmp"
	"errors"
	"slices"
)

type DB struct {
	id    int
	tasks map[int]models.Task
}

func (db *DB) IncrementID() {
	db.id++
}

func (db *DB) GetAllTaks() []models.Task {
	allTasks := []models.Task{}

	for _, tasks := range db.tasks {
		allTasks = append(allTasks, tasks)
	}

	slices.SortFunc(allTasks, func(a, b models.Task) int {
		return cmp.Compare(a.ID, b.ID)
	})

	return allTasks
}

func (db *DB) GetTask(id int) (*models.Task, error) {
	task, exist := db.tasks[id]

	if exist {
		return &task, nil
	}

	return nil, errors.New("Tarefa não encontrada!!!")
}

func NewDB() *DB {
	return &DB{
		id:    0,
		tasks: map[int]models.Task{},
	}
}
