package config

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Log *slog.Logger
	Env Env
	Mux *http.ServeMux
}

type Env struct {
	PORT_SERVER string
}

func (c *Config) InitLogger() {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	c.Log = slog.New(handler)
}

func (c *Config) InitEnv() {
	if err := godotenv.Load(); err != nil {
		c.Log.Error("❌​ Erro ao carregar o arquivo .env", "error", err)
		os.Exit(1)
	}

	c.Env.PORT_SERVER = os.Getenv("PORT_SERVER")
}

func (c *Config) InitMUX() {
	c.Mux = http.NewServeMux()
}
