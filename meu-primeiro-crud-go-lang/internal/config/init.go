package config

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Log      *slog.Logger
	Env      Env
	Mux      *http.ServeMux
	ClientDB *mongo.Client
}

type Env struct {
	SERVER_PORT         string
	MONGO_DB_NAME       string
	MONGO_COLLECTION    string
	MONGO_ROOT_USER     string
	MONGO_ROOT_PASSWORD string
	MONGO_PORT          string
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

	c.Env.SERVER_PORT = os.Getenv("SERVER_PORT")
	c.Env.MONGO_DB_NAME = os.Getenv("MONGO_DB_NAME")
	c.Env.MONGO_COLLECTION = os.Getenv("MONGO_COLLECTION")
	c.Env.MONGO_PORT = os.Getenv("MONGO_PORT")
	c.Env.MONGO_ROOT_PASSWORD = os.Getenv("MONGO_ROOT_PASSWORD")
	c.Env.MONGO_ROOT_USER = os.Getenv("MONGO_ROOT_USER")
}

func (c *Config) InitMUX() {
	c.Mux = http.NewServeMux()
}

func (c *Config) InitDB() {
	uri := fmt.Sprintf(
		"mongodb://%s:%s@localhost:%s/",
		c.Env.MONGO_ROOT_USER,
		c.Env.MONGO_ROOT_PASSWORD,
		c.Env.MONGO_PORT,
	)

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		c.Log.Error("❌​ Erro ao ligar ao MongoDB.", "error", err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		c.Log.Error("❌​ Não foi possível alcançar o MongoDB.", "error", err)
		os.Exit(1)
	}

	c.Log.Info("✅ Conectado ao banco de dados com sucesso!")

	c.ClientDB = client
}
