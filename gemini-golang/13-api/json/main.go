package main

import (
	"encoding/json"
	"fmt"
)

type Conta struct {
	Titular string  `json:"titular"`
	Saldo   float64 `json:"saldo"`
}

func main() {
	c := Conta{Titular: "Dante", Saldo: 400.50}

	dJson, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Erro: %s", err)
	}

	fmt.Println(dJson)
	fmt.Println(string(dJson))

}
