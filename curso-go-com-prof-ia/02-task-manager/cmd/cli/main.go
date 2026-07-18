package main

import (
	clearterminal "02-task-manager/internal/clear-terminal"
	"02-task-manager/internal/db"
	"02-task-manager/internal/task/services"
	"fmt"
)

func main() {
	db := db.NewDB()
	services := services.NewServices(*db)

	for {
		clearterminal.Clear()
		fmt.Println("Digite uma das opções:")

		services.CreateTask()
	}

}
