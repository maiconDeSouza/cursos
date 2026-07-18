package clearterminal

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() error {
	var cmd *exec.Cmd

	// Detecta o sistema operacional rodando o programa
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	// Vincula a saída do comando diretamente ao terminal atual
	cmd.Stdout = os.Stdout

	// Executa o comando
	return cmd.Run()
}
