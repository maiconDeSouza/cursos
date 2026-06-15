package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Conta struct {
	Titular string  `json:"titular"`
	Saldo   float64 `json:"saldo"`
}

type Cliente struct {
	clientes []*Conta
	gate     sync.RWMutex
}

func (cl *Cliente) addConta(co *Conta) {
	cl.gate.Lock()
	defer cl.gate.Unlock()
	cl.clientes = append(cl.clientes, co)
}

func (cl *Cliente) listarContas() []*Conta {
	cl.gate.RLock()
	defer cl.gate.RUnlock()
	return cl.clientes
}

func (cl *Cliente) getConta(titular string) *Conta {
	cl.gate.RLock()
	defer cl.gate.RUnlock()
	contas := cl.clientes
	var conta *Conta

	for _, c := range contas {
		if c.Titular == titular {
			conta = c
		}
	}
	return conta
}

var cl = Cliente{}

func buscarConta(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	getName := r.URL.Query().Get("titular")

	if getName != "" {
		conta := cl.getConta(getName)
		if conta == nil {
			http.Error(w, "Conta não encontrada", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(conta)
		return

	}

	c := cl.listarContas()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(c)
}

func criarConta(w http.ResponseWriter, r *http.Request) {
	var novaConta Conta

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

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
