package main

import (
	"errors"
	"fmt"
)

type Repositorio interface {
	Salvar(dado string) error
}

func GravarDado(r Repositorio, dado string) error {
	err := r.Salvar(dado)

	if err != nil {
		fmt.Printf("Erro ao salvar: %s\n", err)
		return err
	}
	fmt.Println("Dado salvo com sucesso!")
	return err
}

type PostegresDB struct{}
type MockFalho struct{}

func (p PostegresDB) Salvar(dado string) error {
	fmt.Printf("Inserindo no Postgres: %s", dado)
	return nil
}

func (m MockFalho) Salvar(dado string) error {
	err := errors.New("falha de conexão com o banco")

	return err
}

func main() {
	p := PostegresDB{}
	m := MockFalho{}

	err1 := GravarDado(p, "Usuario: Maria")
	if err1 != nil {
		fmt.Println(err1)
	}

	err2 := GravarDado(m, "Usuario: João")
	if err2 != nil {
		fmt.Println(err2)
	}
}
