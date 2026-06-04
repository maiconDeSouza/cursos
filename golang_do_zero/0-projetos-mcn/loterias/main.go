package main

import (
	"bufio"
	"errors"
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
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func sleep(s uint) {
	time.Sleep(time.Duration(s) * time.Second)
}

func responseGame(op string, r *string) error {
	var err error

	switch op {
	case "1":
		megaSena := Games{name: "MegaSena", total: 60, mark: 6}
		*r = fmt.Sprintf("Os números da Sorte da Mega-sena %v\n", megaSena.raffle())
	case "2":
		lotofacil := Games{name: "Lotofácil", total: 25, mark: 15}
		*r = fmt.Sprintf("Os números da Sorte da Lotofácil %v\n", lotofacil.raffle())
	case "3":
		quina := Games{name: "Quina", total: 80, mark: 5}
		*r = fmt.Sprintf("Os números da Sorte da Quina %v\n", quina.raffle())
	case "4":
		lotomania := Games{name: "Lotomania", total: 100, mark: 50}
		*r = fmt.Sprintf("Os números da Sorte da Lotomania %v\n", lotomania.raffle())
	default:
		*r = fmt.Sprintln("Opção Invalida")
		err = errors.New("Opção Invalida")

	}
	return err
}

func main() {
	var r string
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
			sleep(3)
			break
		}

		if err := responseGame(op, &r); err != nil {
			fmt.Println(r)
			fmt.Println("Tente novamente")
			sleep(3)
			continue
		}

		fmt.Println("Aguarde um pouco o resultado...")
		sleep(10)
		clear()
		fmt.Print(r)
		break
	}

}
