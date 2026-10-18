package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	SERVER_PORT            string
	MONGO_DB_NAME          string
	MONGO_COLLECTION_USERS string
	MONGO_COLLECTION_LINKS string
	MONGO_ROOT_USER        string
	MONGO_ROOT_PASSWORD    string
	MONGO_PORT             string
	JWT_SECRET_KEY         string
}

func InitENV() *ENV {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌​ Erro ao carregar o arquivo .env:", err)
	}
	return &ENV{
		SERVER_PORT:            os.Getenv("SERVER_PORT"),
		MONGO_DB_NAME:          os.Getenv("MONGO_DB_NAME"),
		MONGO_COLLECTION_USERS: os.Getenv("MONGO_COLLECTION_USERS"),
		MONGO_COLLECTION_LINKS: os.Getenv("MONGO_COLLECTION_LINKS"),
		MONGO_ROOT_USER:        os.Getenv("MONGO_ROOT_USER"),
		MONGO_ROOT_PASSWORD:    os.Getenv("MONGO_ROOT_PASSWORD"),
		MONGO_PORT:             os.Getenv("MONGO_PORT"),
		JWT_SECRET_KEY:         os.Getenv("JWT_SECRET_KEY"),
	}
}
