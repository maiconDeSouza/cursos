package main

import (
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/errapp"
	"meu-primeiro-crud-go-lang-v2/internal/handler"
	"meu-primeiro-crud-go-lang-v2/internal/infra"
	"meu-primeiro-crud-go-lang-v2/internal/repository"
	"meu-primeiro-crud-go-lang-v2/internal/server"
	"meu-primeiro-crud-go-lang-v2/internal/service"
)

func main() {
	env := config.NewEnv()
	errapp := errapp.NewErrApp()

	db := infra.NewConnectMongo(env)
	repo := repository.NewRepo(db, errapp)

	userService := service.NewUserService(repo, env, errapp)
	adminService := service.NewAdminService(repo, env, errapp)

	userHandler := handler.NewUserHandler(userService, env, errapp)
	adminHandler := handler.NewAdminHandler(adminService, env, errapp)

	server := server.NewServer(env, userHandler, adminHandler)
	server.Run()
}
