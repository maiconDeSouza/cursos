package main

import "fmt"

func recoverEx() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func isApproved(grade uint8) bool {
	defer recoverEx()
	if grade > 10 {
		defer fmt.Println("A nota tem que ser entre 0 e 10")
		panic("A nota tem que ser entre 0 e 10")
	}

	defer fmt.Println("O Aluno está aprovado?")
	if grade >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isApproved(12))
	fmt.Println("O programa continuou...")
}
