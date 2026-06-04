package main

import "fmt"

func isApproved(grade uint8) bool {
	if grade > 10 {
		fmt.Println("A nota tem que ser entre 0 e 10")
		return false
	}

	defer fmt.Println("O Aluno está aprovado?")
	if grade >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isApproved(8))
	fmt.Println(isApproved(12))
	fmt.Println(isApproved(4))

}
