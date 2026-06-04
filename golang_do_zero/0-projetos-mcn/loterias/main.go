package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
	"time"
)

type Games struct {
	name  string
	total uint
	mark  uint
}

func (g Games) generate() []uint {
	var numbers []uint

	for i := 1; i <= int(g.total); i++ {
		numbers = append(numbers, uint(i))
	}

	return numbers
}

func (g Games) raffle() []uint {
	var r []uint
	n := g.generate()

	for i := 1; i <= int(g.mark); i++ {
		index := rand.Intn(len(n))
		r = append(r, n[index])
		n = slices.Delete(n, index, index+1)
	}

	slices.Sort(r)
	return r
}

func clear() {
	var cmd *exec.Cmd

	// Detecta o sistema operacional atual
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Funciona para Linux, macOS, Unix, etc.
		cmd = exec.Command("clear")
	}

	// Vincula a saída do comando direto no terminal atual
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func responseGame(op string) (string, bool) {
	megaSena := Games{name: "MegaSena", total: 60, mark: 6}
	lotofacil := Games{name: "Lotofácil", total: 25, mark: 15}
	quina := Games{name: "Quina", total: 80, mark: 5}
	lotomania := Games{name: "Lotomania", total: 100, mark: 50}

	var response string
	e := false

	switch op {
	case "1":
		response = fmt.Sprintf("Os números da Sorte da Mega-sena %v\n", megaSena.raffle())
	case "2":
		response = fmt.Sprintf("Os números da Sorte da Lotofácil %v\n", lotofacil.raffle())
	case "3":
		response = fmt.Sprintf("Os números da Sorte da Quina %v\n", quina.raffle())
	case "4":
		response = fmt.Sprintf("Os números da Sorte da Lotomania %v\n", lotomania.raffle())
	default:
		response = fmt.Sprintln("Opção Invalida")
		e = true
	}
	return response, e
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clear()
		fmt.Println("Digite uma das opções:")
		fmt.Println("1) Mega-Sena")
		fmt.Println("2) Lotofácil")
		fmt.Println("3) Quina")
		fmt.Println("4) Lotomania")
		fmt.Println("5) Sair")

		op, err := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		if err != nil {
			fmt.Println("Erro ao ler o prompt:", err)
			return
		}

		clear()

		if op == "5" {
			fmt.Println("Tchauuuuu!!!")
			time.Sleep(time.Second)
			break
		}

		r, e := responseGame(op)

		if e {
			fmt.Println(r)
			fmt.Println("Tente novamente")
			time.Sleep(time.Second)
			continue
		}

		fmt.Print(r)
		break
	}

}
