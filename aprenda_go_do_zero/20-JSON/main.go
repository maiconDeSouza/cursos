package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Dog struct {
	Name  string `json:"name"`
	Age   uint   `json:"age"`
	Breed string `json:"breed"`
}

func main() {
	c := Dog{"Dona Maia", 11, "Lavrador"}
	c2 := Dog{}

	fmt.Println("Objeto original:", c)

	cJSON, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON gerado:", bytes.NewBuffer(cJSON))

	c2JSON, err := os.ReadFile("dog.json")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	fmt.Println("Conteúdo lido do dog.json:", string(c2JSON))

	err = json.Unmarshal(c2JSON, &c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Objeto preenchido pelo arquivo:", c2)
}
