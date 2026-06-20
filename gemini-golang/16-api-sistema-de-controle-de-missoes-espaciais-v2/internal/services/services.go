package services

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/models"
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/storage"
	"encoding/json"
	"errors"
	"io"
)

var st = storage.Storage{}

func CreateShip(body io.ReadCloser) (*models.Ship, error) {
	sh := models.Ship{}
	err := json.NewDecoder(body).Decode(&sh)
	if err != nil {
		return nil, errors.New("Erro ao ler o JSON")
	}

	st.AddShip(&sh)

	return &sh, nil
}

func GetAllShips() []*models.Ship {
	return st.GetAllShips()
}

func GetAllActives() []*models.Ship {
	return st.GetAllActives()
}

func EditShip(name string, body io.ReadCloser) (*models.Ship, error) {
	newSH := models.Ship{}
	err := json.NewDecoder(body).Decode(&newSH)
	if err != nil {
		return nil, errors.New("Erro ao ler o JSON")
	}

	sh, exist := st.UpdateShip(name, &newSH)
	if !exist {
		return nil, errors.New("Nave não encontrada!")
	}

	return sh, nil
}

func DeleteShip(name string) (*models.Ship, error) {
	sh, exist := st.DeleteShip(name)

	if !exist {
		return nil, errors.New("Nave não encontrada!")
	}

	return sh, nil
}
