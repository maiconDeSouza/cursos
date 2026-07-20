package main

import (
	handlersIP "api-busca-ip-serve/internal/ip/handlersIP"
	servicesIP "api-busca-ip-serve/internal/ip/services"
	"api-busca-ip-serve/internal/router"
	handlersServer "api-busca-ip-serve/internal/server/handlersServer"
	servicesServer "api-busca-ip-serve/internal/server/services"
	"fmt"
	"log"
	"net/http"
)

const PORT = 2005

func main() {
	servicesIP := servicesIP.NewSerice()
	handlersIP := handlersIP.NewHandlers(servicesIP)

	servicesServices := servicesServer.NewSerice()
	handlersServer := handlersServer.NewHandlers(servicesServices)

	mux := router.InitRoutes(handlersIP, handlersServer)

	fmt.Printf("Servidor rodando na porta :%d", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
