package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	SERVER_PORT         string
	MONGO_DB_NAME       string
	MONGO_COLLECTION    string
	MONGO_ROOT_USER     string
	MONGO_ROOT_PASSWORD string
	MONGO_PORT          string
}

func NewEnv() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌​ Erro ao carregar o arquivo .env:", err)
	}

	return &Env{
		SERVER_PORT:         os.Getenv("SERVER_PORT"),
		MONGO_DB_NAME:       os.Getenv("MONGO_DB_NAME"),
		MONGO_COLLECTION:    os.Getenv("MONGO_COLLECTION"),
		MONGO_ROOT_USER:     os.Getenv("MONGO_ROOT_USER"),
		MONGO_ROOT_PASSWORD: os.Getenv("MONGO_ROOT_PASSWORD"),
		MONGO_PORT:          os.Getenv("MONGO_PORT"),
	}
}
