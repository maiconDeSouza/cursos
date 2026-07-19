package router

import (
	handlersIP "api-busca-ip-serve/internal/ip/handlers"
	handlersServer "api-busca-ip-serve/internal/server/handlers"
	"net/http"
)

func InitRoutes(handlersIP *handlersIP.HandlersIP, handlersServer *handlersServer.HandlersServer) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/ip", handlersIP.GetIP)
	mux.HandleFunc("GET /api/v1/server", handlersServer.GetServer)

	return mux
}
