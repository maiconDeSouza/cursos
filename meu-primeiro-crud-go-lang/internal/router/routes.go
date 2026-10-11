package router

import (
	"fmt"
	"meu-primeiro-crud-go-lang/internal/config"
	"meu-primeiro-crud-go-lang/internal/handlers"
	"net/http"
)

type Routes struct {
	handlers handlers.HandlersInterface
	cfg      *config.Config
}

func NewRoutes(mux *http.ServeMux, handlers handlers.HandlersInterface, cfg *config.Config) *Routes {
	return &Routes{handlers: handlers, cfg: cfg}
}

func (r *Routes) InitRoutes() {
	r.cfg.Mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá")
	})
	r.cfg.Mux.HandleFunc("POST /user", r.handlers.CreateUser)
}
