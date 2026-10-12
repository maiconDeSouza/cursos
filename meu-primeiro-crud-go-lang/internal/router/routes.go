package router

import (
	"fmt"
	"meu-primeiro-crud-go-lang/internal/config"
	"meu-primeiro-crud-go-lang/internal/handlers"
	"meu-primeiro-crud-go-lang/internal/middleware"
	"net/http"
)

type Routes struct {
	handlers   handlers.HandlersInterface
	middleware middleware.Middleware
	cfg        *config.Config
}

func NewRoutes(mux *http.ServeMux, handlers handlers.HandlersInterface, middleware middleware.Middleware, cfg *config.Config) *Routes {
	return &Routes{handlers: handlers, middleware: middleware, cfg: cfg}
}

func (r *Routes) InitRoutes() {
	r.cfg.Mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá")
	})
	r.cfg.Mux.HandleFunc("POST /user", r.handlers.CreateUser)
	r.cfg.Mux.HandleFunc("POST /user/login", r.handlers.LoginUser)
	r.cfg.Mux.HandleFunc("PUT /user/update", r.middleware.AuthMiddleware(r.handlers.UpdateUser))
	r.cfg.Mux.HandleFunc("DELETE /user/delete", r.middleware.AuthMiddleware(r.handlers.DeleteUser))

}
