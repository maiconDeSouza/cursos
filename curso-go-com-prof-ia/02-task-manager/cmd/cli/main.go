package main

import (
	"02-task-manager/internal/tasks/repositories"
	"02-task-manager/internal/tasks/services"
	"02-task-manager/internal/terminal"
	"fmt"
)

func main() {
	db := repositories.NewDB()
	services := services.NewServices(*db)

	for {
		terminal.Clear()
		fmt.Println("Selecione uma das opções")
		fmt.Println()
		terminal.WriteOptions()
		op, err := terminal.ScanText()
		if err != nil {
			terminal.Clear()
			fmt.Printf("Erro: %t", err)
			terminal.Time(5)
			continue
		}

		switch op {
		case "1":
			services.AddTasks()
		case "2":
			services.DoneTask()
		case "3":
			services.GetAllTasks()
		case "4":
			services.GetTask()
		case "5":
			fmt.Println("Até logo")
			terminal.Time(3)
			return
		default:
			fmt.Println("Opção invalida")
			terminal.Time(3)
		}

	}

}
