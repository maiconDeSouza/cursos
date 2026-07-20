package router

import (
	handlersIP "api-busca-ip-serve/internal/ip/handlersIP"
	handlersServer "api-busca-ip-serve/internal/server/handlersServer"
	"net/http"
)

func InitRoutes(handlersIP *handlersIP.HandlersIP, handlersServer *handlersServer.HandlersServer) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/v1/ip", handlersIP.GetIP)
	mux.HandleFunc("POST /api/v1/server", handlersServer.GetServer)

	return mux
}
