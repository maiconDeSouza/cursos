package services

import (
	"02-task-manager/internal/str"
	"02-task-manager/internal/tasks/interfaces"
	"02-task-manager/internal/tasks/models"
	"02-task-manager/internal/tasks/repositories"
	"02-task-manager/internal/terminal"
	"fmt"
	"strconv"
	"strings"
)

type Services struct {
	repo interfaces.DB
}

func (s Services) AddTasks() {
	terminal.Clear()
	fmt.Println("Escreve uma ou mais Tarefas separadas por virgulas:")
	tasksString, err := terminal.ScanText()
	if err != nil {
		terminal.Clear()
		fmt.Printf("Erro: %t", err)
		terminal.Time(5)
		return
	}

	tasksSlice := str.StringToSlice(tasksString)
	tasksSalve := []models.Task{}

	for _, text := range tasksSlice {
		task := models.Task{
			Title: text,
			Done:  false,
		}
		newTask := s.repo.Salve(&task)
		tasksSalve = append(tasksSalve, newTask)
	}
	terminal.RenderTasks(tasksSalve)
	terminal.Continue()
}

func (s Services) GetAllTasks() {
	tasks := s.repo.GetAllTaks()

	terminal.RenderTasks(tasks)
	terminal.Continue()
}

func (s Services) GetTask() {
	tasks := s.repo.GetAllTaks()
	terminal.RenderTasks(tasks)

	fmt.Println()
	fmt.Println("Digite o ID da tasks que você quer ver ou Q para sair")
	r, err := terminal.ScanText()
	if err != nil {
		terminal.Clear()
		fmt.Printf("Erro: %t", err)
		terminal.Time(5)
		return
	}

	if strings.ToLower(r) == "q" {
		fmt.Println("Até logo!!!")
		terminal.Time(3)
		return
	}

	n, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println("Apenas digite números ou q")
		terminal.Time(3)
		return
	}

	task, err := s.repo.GetTask(n)
	if err != nil {
		fmt.Printf("Erro: %s", err)
		terminal.Time(3)
		return
	}

	terminal.RenderTasks([]models.Task{*task})
	terminal.Continue()
}

func (s Services) DoneTask() {
	tasks := s.repo.GetAllTaks()
	terminal.RenderTasks(tasks)

	fmt.Println()
	fmt.Println("Digite o ID da tasks que você concluiu ou Q para sair")
	r, err := terminal.ScanText()
	if err != nil {
		terminal.Clear()
		fmt.Printf("Erro: %t", err)
		terminal.Time(5)
		return
	}

	if strings.ToLower(r) == "q" {
		fmt.Println("Até logo!!!")
		terminal.Time(3)
		return
	}

	n, err := strconv.Atoi(r)
	if err != nil {
		fmt.Println("Apenas digite números ou q")
		terminal.Time(3)
		return
	}

	err = s.repo.DoneTask(n)
	if err != nil {
		fmt.Printf("Erro: %s", err)
		terminal.Time(3)
		return
	}

	tasks = s.repo.GetAllTaks()

	terminal.RenderTasks(tasks)
	terminal.Continue()

}

func NewServices(db repositories.DB) *Services {
	return &Services{
		repo: &db,
	}
}
