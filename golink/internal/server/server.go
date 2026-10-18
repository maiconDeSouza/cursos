package server

import (
	"fmt"
	"golink/internal/config"
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
	env *config.ENV
}

func NewServer(env *config.ENV) *Server {
	mux := http.NewServeMux()
	return &Server{mux: mux, env: env}
}

func (s *Server) Run() {
	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK:OK")
	})

	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + s.env.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.env.SERVER_PORT), s.mux))
}
