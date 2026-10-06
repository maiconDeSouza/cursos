package main

import (
	"fmt"
	"log"
	"meu-primeiro-crud-go-lang/src/controller/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TEST"))

	mux := http.NewServeMux()

	routes.InitRoutes(mux)
	fmt.Println("🚀 Servidor iniciado com sucesso na porta ", 1992)
	if err := http.ListenAndServe(":1992", mux); err != nil {
		log.Fatal(err)
	}
}
