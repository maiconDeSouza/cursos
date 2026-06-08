package main

import (
	"fmt"
	"sistema-banco/banco"
	"sistema-banco/terminal"
)

func gerencia() {
	for {
		terminal.NomeBanco()
		fmt.Println("Digite a sua senha de gerente:")
		senhaGerente := terminal.ScanTerminal()

		err := banco.Gerencia.LoginGerencia(senhaGerente)
		if err != nil {
			fmt.Printf("Erro: %s", err)
			terminal.Sleep(2)
			continue
		}

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
			_, conta := banco.CriarConta(nome, tipo, senha)
			fmt.Printf("Conta de número %s aberta com sucesso!", conta)
			terminal.Sleep(5)
		case "2":
			terminal.NomeBanco()
			fmt.Println("Abrindo conta:")
			banco.TodasContas()
			fmt.Println()
			fmt.Println("Aperte qualquer tecla para sair")
			sair := terminal.ScanTerminal()
			if sair != "" {
				continue
			}
			continue
		case "3":
			fmt.Println("Até logo...")
			terminal.Sleep(3)
			return
		default:
			fmt.Println("Valor errado!")
			continue
		}
	}
}

func clientes() {
	for {
		terminal.NomeBanco()
		fmt.Println("1 - Sacar")
		fmt.Println("2 - Depositar")
		fmt.Println("3 - Saldo")
		fmt.Println()
		fmt.Println("Digite uma das opções ou apenas ENTER para sair")
		op := terminal.ScanTerminal()

		switch op {
		case "":
			fmt.Println("Até logo...")
			terminal.Sleep(3)
			return
		case "1":
			terminal.NomeBanco()
			fmt.Println("Digite o número da sua conta ou apenas ENTER para sair:")
			fmt.Println()
			numeroDaConta := terminal.ScanTerminal()
			if numeroDaConta == "" {
				terminal.Sleep(3)
				continue
			}
			conta, errConta := banco.BuscarConta(numeroDaConta)
			if errConta != nil {
				fmt.Printf("Erro: %s", errConta)
				terminal.Sleep(3)
				continue
			}
			// escape := false
			for {
				terminal.NomeBanco()
				fmt.Println("digite sua senha ou apenas ENTER para sair:")
				senha := terminal.ScanTerminal()
				if senha == "" {
					break
				}
				errSenha := banco.VerificarSenha(conta, senha)
				if errSenha != nil {
					fmt.Printf("Erro: %s", errSenha)
					terminal.Sleep(3)
					continue
				}

				terminal.NomeBanco()
				fmt.Println("digite o valor do saque ou apenas ENTER para sair:")
				valor := terminal.ScanTerminal()
				if valor == "" {
					// escape = true
					break
				}
				valorFloat, errValor := terminal.ConverterValor(valor)
				if errValor != nil {
					fmt.Printf("Erro: %s", errValor)
					terminal.Sleep(3)
					continue
				}
				errSaque := banco.OperacaoDeSaque(conta, valorFloat)
				if errSaque != nil {
					fmt.Printf("Erro: %s", errSaque)
					terminal.Sleep(3)
					continue
				}
				break
			}
			// if escape {
			// 	continue
			// }
		case "2":
			terminal.NomeBanco()
			fmt.Println("Digite o número da sua conta ou apenas ENTER para sair:")
			fmt.Println()
			numeroDaConta := terminal.ScanTerminal()
			if numeroDaConta == "" {
				terminal.Sleep(3)
				continue
			}
			conta, errConta := banco.BuscarConta(numeroDaConta)
			if errConta != nil {
				fmt.Printf("Erro: %s", errConta)
				terminal.Sleep(3)
				continue
			}
			for {
				terminal.NomeBanco()
				fmt.Println("digite o valor do Deposito ou apenas ENTER para sair:")
				valor := terminal.ScanTerminal()
				if valor == "" {
					break
				}
				valorFloat, errValor := terminal.ConverterValor(valor)
				if errValor != nil {
					fmt.Printf("Erro: %s", errValor)
					terminal.Sleep(3)
					continue
				}
				errDeposito := banco.OperacaoDeDeposito(conta, valorFloat)
				if errDeposito != nil {
					fmt.Printf("Erro: %s", errDeposito)
					terminal.Sleep(3)
					continue
				}
				break
			}
		case "3":
			terminal.NomeBanco()
			fmt.Println("Digite o número da sua conta ou apenas ENTER para sair:")
			fmt.Println()
			numeroDaConta := terminal.ScanTerminal()
			if numeroDaConta == "" {
				terminal.Sleep(3)
				continue
			}
			conta, errConta := banco.BuscarConta(numeroDaConta)
			if errConta != nil {
				fmt.Printf("Erro: %s", errConta)
				terminal.Sleep(3)
				continue
			}
			for {
				terminal.NomeBanco()
				fmt.Println("digite sua senha ou apenas ENTER para sair:")
				senha := terminal.ScanTerminal()
				if senha == "" {
					break
				}
				errSenha := banco.VerificarSenha(conta, senha)
				if errSenha != nil {
					fmt.Printf("Erro: %s", errSenha)
					terminal.Sleep(3)
					continue
				}

				saldo := banco.OperacaoDeSaldo(conta)

				fmt.Printf("Olá %s seu salto em conta é R$%.2f", conta.PegarNome(), saldo)
				terminal.Sleep(10)
				break
			}
		}

	}

}

func main() {
	terminal.NomeBanco()
	fmt.Println("1 - Gerente")
	fmt.Println("2 - Cliente")
	fmt.Println()
	fmt.Println("Digite 1 ou 2:")
	op := terminal.ScanTerminal()

	if op == "1" {
		gerencia()
	}
	clientes()
}
