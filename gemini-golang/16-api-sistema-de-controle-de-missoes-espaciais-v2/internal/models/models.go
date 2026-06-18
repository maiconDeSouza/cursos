package models

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/storage"
	"encoding/json"
	"errors"
	"io"
)

type Ship struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Active bool   `json:"active"`
}

var st = storage.Storage{}

func (sh *Ship) CreateShip(body io.ReadCloser) error {
	err := json.NewDecoder(body).Decode(sh)
	if err != nil {
		return errors.New("Erro ao ler o JSON")
	}
	st.AddShip(sh)

	return nil
}
