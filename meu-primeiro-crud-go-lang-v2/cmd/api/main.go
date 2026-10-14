package main

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/handler"
	"meu-primeiro-crud-go-lang-v2/internal/infra"
	"meu-primeiro-crud-go-lang-v2/internal/repository"
	"meu-primeiro-crud-go-lang-v2/internal/server"
	"meu-primeiro-crud-go-lang-v2/internal/service"
)

func main() {
	env := config.NewEnv()
	db := infra.NewConnectMongo(env)
	repo := repository.NewRepo(db)

	userService := service.NewUserService(repo, env)
	adminService := service.NewAdminService(repo, env)

	userHandler := handler.NewUserHandler(userService, env)
	adminHandler := handler.NewAdminHandler(adminService, env)

	server := server.NewServer(env, userHandler, adminHandler)
	server.Run()
}
