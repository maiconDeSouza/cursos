package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

type Account struct {
	Holder  string  `json:"holder"`
	Balance float64 `json:"balance"`
}

type Customer struct {
	customers []*Account
	gate      sync.RWMutex
}

func (c *Customer) addCustomer(a *Account) {
	c.customers = append(c.customers, a)
}

var c Customer

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c.gate.RLock()
	defer c.gate.RUnlock()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(c.customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	c.gate.RLock()
	defer c.gate.RUnlock()
	holderName := r.URL.Query().Get("holder")

	for _, c := range c.customers {
		if c.Holder == holderName {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(c)
			return
		}
	}
	http.Error(w, "Conta não encontrada", http.StatusNotFound)
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	c.gate.Lock()
	defer c.gate.Unlock()
	a := Account{}

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, "Erro ao ler o JSON", http.StatusBadRequest)
		return
	}

	c.addCustomer(&a)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)

	customersJson, _ := json.MarshalIndent(c.customers, "", " ")

	err = os.WriteFile("accounts.json", customersJson, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err.Error())
	}
}

func requestMethods(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	holder := r.URL.Query().Get("holder")

	switch {
	case method == http.MethodGet && holder != "":
		getCustomer(w, r)
	case method == http.MethodGet:
		getAllCustomers(w, r)
	case method == http.MethodPost:
		createAccount(w, r)
	default:
		http.Error(w, "Contas não encontrada", http.StatusNotFound)
	}

}

func init() {
	j, err := os.ReadFile("accounts.json")

	if err != nil {
		panic("Erro ao ler o JSON: " + err.Error())
	}

	err = json.Unmarshal(j, &c.customers)

	if err != nil {
		panic("Erro ao converter o JSON para a struct: " + err.Error())
	}
}

func main() {
	http.HandleFunc("/account", requestMethods)
	fmt.Println("Servidor rodando!")
	http.ListenAndServe(":2005", nil)
}
