package main

import (
	"meu-primeiro-crud-go-lang/internal/config"
	"meu-primeiro-crud-go-lang/internal/handlers"
	"meu-primeiro-crud-go-lang/internal/repositories"
	"meu-primeiro-crud-go-lang/internal/router"
	"meu-primeiro-crud-go-lang/internal/services"
	"net/http"
	"os"
)

func main() {
	cfg := config.Config{}
	cfg.InitLogger()
	cfg.InitEnv()
	cfg.InitMUX()
	cfg.InitDB()

	repo := repositories.NewRepositories(cfg.ClientDB, &cfg)
	services := services.NewServices(repo, &cfg)
	handlers := handlers.NewHandlers(services, &cfg)
	routes := router.NewRoutes(cfg.Mux, handlers, &cfg)
	routes.InitRoutes()

	cfg.Log.Info("🚀 Servidor iniciado com sucesso na porta", "port", cfg.Env.SERVER_PORT)
	if err := http.ListenAndServe(":"+cfg.Env.SERVER_PORT, cfg.Mux); err != nil {
		cfg.Log.Error("Erro ao iniciar o servidor", "error", err)
		os.Exit(1)
	}
}
