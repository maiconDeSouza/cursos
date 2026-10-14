package server

import (
	"fmt"
	"log"
	"meu-primeiro-crud-go-lang-v2/internal/config"
	"meu-primeiro-crud-go-lang-v2/internal/handler"
	"net/http"
)

type Server struct {
	env          *config.Env
	userHandler  handler.UserHandlerInterface
	adminHandler handler.AdminHandlerInterface
}

func NewServer(env *config.Env, userHandler handler.UserHandlerInterface, adminHandler handler.AdminHandlerInterface) *Server {
	return &Server{env: env, userHandler: userHandler, adminHandler: adminHandler}
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "server top")
	})

	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + s.env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.env.SERVER_PORT), mux))
}
