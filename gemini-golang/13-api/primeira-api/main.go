package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Conta struct {
	Titular string  `json:"titular"`
	Saldo   float64 `json:"saldo"`
}

type Cliente struct {
	clientes []*Conta
}

func (cl *Cliente) addConta(co *Conta) {
	cl.clientes = append(cl.clientes, co)
}

func (cl *Cliente) listarContas() []*Conta {
	return cl.clientes
}

var cl = Cliente{}

func buscarConta(w http.ResponseWriter, r *http.Request) {
	c := cl.listarContas()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(c)
}

func criarConta(w http.ResponseWriter, r *http.Request) {
	var novaConta Conta

	err := json.NewDecoder(r.Body).Decode(&novaConta)

	if err != nil {
		http.Error(w, "Erro ao ler o JSON", http.StatusBadRequest)
		return
	}

	cl.addConta(&novaConta)

	fmt.Printf("Nova conta recebida! Titular: %s | Saldo: %.2f\n", novaConta.Titular, novaConta.Saldo)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Conta criada com sucesso!"))
}

func main() {

	http.HandleFunc("/buscar", buscarConta)
	http.HandleFunc("/criar", criarConta)
	fmt.Println("Servidor rodando...")
	http.ListenAndServe(":2005", nil)
}
