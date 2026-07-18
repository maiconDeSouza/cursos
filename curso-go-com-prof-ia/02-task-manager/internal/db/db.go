package db

import "02-task-manager/internal/task/models"

type DB struct {
	Index uint
	Tasks map[uint]models.Task
}

func NewDB() *DB {
	db := &DB{
		Index: 0,
		Tasks: make(map[uint]models.Task),
	}
	return db
}
