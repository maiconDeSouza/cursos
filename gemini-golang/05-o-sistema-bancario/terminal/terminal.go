package terminal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func Sleep(s uint) {
	time.Sleep(time.Duration(s) * time.Second)
}

func PrintTrace() {
	fmt.Println("-----------------------------------------")
}

func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ScanTerminal() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	op := scanner.Text()

	return op
}

func NomeBanco() {
	Clear()
	fmt.Println("Sistema Bancario")
	PrintTrace()
	PrintTrace()
	PrintTrace()
	fmt.Println()
}

func ConverterValor(valor string) (float64, error) {
	numero, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		errSistema := errors.New("Você digitou o valor em forma errado. Tenhe digitar algo assim 250.00")
		return 0.00, errSistema
	}
	return numero, nil
}
