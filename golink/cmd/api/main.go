package main

import (
	"golink/internal/config"
	"golink/internal/database"
	"golink/internal/server"
)

func main() {
	env := config.InitENV()
	_ = database.NewConnectMongo(env)
	server := server.NewServer(env)
	server.Run()
}
