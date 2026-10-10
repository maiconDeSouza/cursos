package main

import (
	"fmt"
	"meu-primeiro-crud-go-lang/internal/config"
	"net/http"
	"os"
)

func main() {
	cfg := config.Config{}
	cfg.InitLogger()
	cfg.InitEnv()
	cfg.InitMUX()

	cfg.Mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "oi")
	})

	cfg.Log.Info("🚀 Servidor iniciado com sucesso na porta", "port", cfg.Env.PORT_SERVER)
	if err := http.ListenAndServe(":"+cfg.Env.PORT_SERVER, cfg.Mux); err != nil {
		cfg.Log.Error("Erro ao iniciar o servidor", "error", err)
		os.Exit(1)
	}
}
