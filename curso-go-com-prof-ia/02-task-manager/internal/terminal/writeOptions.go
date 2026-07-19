package terminal

import "fmt"

var options = []string{
	"1 – Adicionar tarefa",
	"2 – Concluir tarefa",
	"3 – Listar todas",
	"4 – Buscar por ID",
	"5 – Sair",
}

func WriteOptions() {
	for _, op := range options {
		fmt.Println(op)
	}
}
