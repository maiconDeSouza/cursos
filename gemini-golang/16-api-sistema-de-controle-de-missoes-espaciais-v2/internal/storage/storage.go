package storage

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/models"
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
