package main

import (
	"golink/internal/config"
	"golink/internal/database"
	"golink/internal/handler"
	"golink/internal/middleware"
	"golink/internal/repository"
	"golink/internal/server"
	"golink/internal/service"
)

func main() {
	env := config.InitENV()
	collection := database.NewConnectMongo(env)

	userRepo := repository.NewUserRepository(collection, env)
	linkRepo := repository.NewLinkRepository(collection, env)

	userService := service.NewUserService(userRepo, env)
	linkService := service.NewLinkService(linkRepo, env)

	userHandler := handler.NewUserHandler(userService, env)
	linkHandler := handler.NewLinkHandler(linkService, env)

	middleware := middleware.NewMiddleware(env)

	server := server.NewServer(userHandler, linkHandler, middleware, env)
	server.Run()
}
