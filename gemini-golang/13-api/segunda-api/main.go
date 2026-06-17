package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
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

func (c *Customer) accountExists(holder string) (int, *Account, error) {
	for i, c := range c.customers {
		if c.Holder == holder {
			return i, c, nil
		}
	}

	return -1, nil, errors.New("Cliente não encontrado!")
}

func (c *Customer) deleteAccount(i int) {
	c.customers = slices.Delete(c.customers, i, i+1)
}

var c Customer

func wJson() {
	customersJson, _ := json.MarshalIndent(c.customers, "", " ")

	err := os.WriteFile("accounts.json", customersJson, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar arquivo:", err.Error())
	}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c.gate.RLock()
	defer c.gate.RUnlock()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(c.customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	c.gate.RLock()
	defer c.gate.RUnlock()
	holderName := r.URL.Query().Get("holder")

	_, c, err := c.accountExists(holderName)
	if err != nil {
		http.Error(w, "Conta não encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(c)
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

	wJson()
}

func editAccount(w http.ResponseWriter, r *http.Request) {
	c.gate.Lock()
	defer c.gate.Unlock()
	holderName := r.URL.Query().Get("holder")
	a := Account{}

	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Erro ao ler o JSON", http.StatusBadRequest)
		return
	}

	if holderName == "" {
		http.Error(w, "Conta não encontrada", http.StatusNotFound)
		return
	}

	_, c, err := c.accountExists(holderName)
	if err != nil {
		http.Error(w, "Conta não encontrada", http.StatusNotFound)
		return
	}

	c.Holder = a.Holder
	c.Balance = a.Balance

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
	wJson()
}

func delAccount(w http.ResponseWriter, r *http.Request) {
	c.gate.Lock()
	defer c.gate.Unlock()
	holderName := r.URL.Query().Get("holder")

	if holderName == "" {
		http.Error(w, "Conta não encontrada", http.StatusNotFound)
		return
	}

	i, a, err := c.accountExists(holderName)
	if err != nil {
		http.Error(w, "Conta não encontrada", http.StatusNotFound)
		return
	}

	c.deleteAccount(i)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
	wJson()
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
	case method == http.MethodPut:
		editAccount(w, r)
	case method == http.MethodDelete:
		delAccount(w, r)
	default:
		http.Error(w, "Contas não encontrada", http.StatusMethodNotAllowed)
	}

}

func init() {
	j, err := os.ReadFile("accounts.json")

	if err != nil {
		fmt.Println("Nenhum arquivo encontrado, iniciando banco de dados vazio...")
		return
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
