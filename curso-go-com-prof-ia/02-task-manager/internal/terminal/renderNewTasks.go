package terminal

import (
	"02-task-manager/internal/tasks/models"
	"fmt"
)

func RenderTasks(newTask []models.Task) {
	Clear()
	fmt.Println("Tarefas criadas")
	fmt.Println()

	for _, task := range newTask {
		fmt.Printf("ID: %d - %s - Done: %t\n", task.ID, task.Title, task.Done)
	}
}
