package main

import (
	"02-task-manager/internal/tasks/repositories"
	"02-task-manager/internal/tasks/services"
)

func main() {
	db := repositories.NewDB()
	services := services.NewServices(*db)
	services.Qualquer()
}
