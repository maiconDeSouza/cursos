package storage

import (
	"api-sistema-de-controle-de-missoes-espaciais/internal/models"
	"errors"
	"sync"
)

type Storage struct {
	ships []*models.Ship
	gate  sync.RWMutex
}

func (st *Storage) AddShip(sh *models.Ship) {
	st.gate.Lock()
	defer st.gate.Unlock()
	st.ships = append(st.ships, sh)
}

func (st *Storage) GetSships() []*models.Ship {
	st.gate.RLock()
	defer st.gate.RUnlock()
	return st.ships
}

func (st *Storage) ShipExists(nameShip string) (*models.Ship, error) {
	st.gate.RLock()
	defer st.gate.Unlock()
	for _, s := range st.ships {
		if s.Name == nameShip {
			return s, nil
		}
	}
	return nil, errors.New("Nave não encontrada!")
}

func (st *Storage) UPship(upName, upStatus string, upActive bool, sh *models.Ship) {
	st.gate.Lock()
	defer st.gate.Unlock()

	if upName != "" {
		sh.Name = upName
	}

	if upStatus != "" {
		sh.Status = upStatus
	}

	if sh.Active != upActive {
		sh.Active = upActive
	}
}

var StoageShips = &Storage{}
