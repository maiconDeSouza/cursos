package main

import (
	"fmt"
	"sistema-banco/banco"
	"sistema-banco/terminal"
)

func gerencia() {
	for {
		terminal.NomeBanco()
		fmt.Println("1 - Abrir uma nova conta")
		fmt.Println("2 - Listar todas contas")
		fmt.Println("3 - Ir para a area de clientes")
		fmt.Println()
		fmt.Println("Digite uma das opções:")
		op2 := terminal.ScanTerminal()

		switch op2 {
		case "1":
			terminal.NomeBanco()
			fmt.Println("Abrindo conta:")
			fmt.Println("Digite cc -> para conta corrente")
			fmt.Println("Digite cp -> para conta poupança")
			tipo := terminal.ScanTerminal()
			fmt.Println()

			fmt.Println("Nome do Cliente:")
			nome := terminal.ScanTerminal()
			fmt.Println()

			var senha string
			for {
				fmt.Println("Digite uma senha númerica de 4 a 6 digitos:")
				senha = terminal.ScanTerminal()

				if len(senha) < 4 || len(senha) > 6 {
					fmt.Println("Senha inválida!!!")
					terminal.Sleep(2)
					continue
				}
				break
			}
			banco.CriarConta(nome, tipo, senha)
		}

	}
}

func clientes() {
	for {
		terminal.NomeBanco()
		fmt.Println("1 - ")

	}
}

func main() {
	terminal.NomeBanco()
	fmt.Println("1 - Gerente")
	fmt.Println("2 - Cliente")
	fmt.Println()
	fmt.Println("Digite 1 ou 2:")
	// op := terminal.ScanTerminal()

	// if op != "1" {
	// 	terminal.Sleep(2)
	// 	return
	// }
	// terminal.Sleep(2)
	// terminal.NomeBanco()
	// fmt.Println("Digite a senha de gerente:")
	// senha := terminal.ScanTerminal()
	// errLogin := banco.Gerencia.LoginGerencia(senha)
	// if errLogin != nil {
	// 	fmt.Printf("Erro: %s", errLogin)
	// 	terminal.Sleep(2)
	// 	continue
	// }

}
