package server

import (
	"fmt"
	"golink/internal/config"
	"golink/internal/handler"
	"golink/internal/middleware"
	"log"
	"net/http"
)

type Server struct {
	userHandler handler.UserHandlerInterface
	linkHandler handler.LinkHandlerInterface
	middleware  middleware.MiddlewareInterface
	mux         *http.ServeMux
	env         *config.ENV
}

func NewServer(
	userHandler handler.UserHandlerInterface,
	linkHandler handler.LinkHandlerInterface,
	middleware middleware.MiddlewareInterface,
	env *config.ENV,
) *Server {
	mux := http.NewServeMux()
	return &Server{userHandler: userHandler, linkHandler: linkHandler, middleware: middleware, mux: mux, env: env}
}

func (s *Server) Run() {
	s.routerHome()
	s.routerUser()
	s.routerLink()

	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + s.env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.env.SERVER_PORT), s.mux))
}

func (s *Server) routerHome() {
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK:OK")
	})

	s.mux.HandleFunc("GET /r/{title}", s.linkHandler.Redirect)
}

func (s *Server) routerUser() {
	s.mux.HandleFunc("POST /user", s.userHandler.CreateUser)
	s.mux.HandleFunc("POST /user/login", s.userHandler.Login)
}

func (s *Server) routerLink() {
	s.mux.HandleFunc("POST /user/auth/link", s.middleware.AuthUser(s.linkHandler.CreateLink))
}
