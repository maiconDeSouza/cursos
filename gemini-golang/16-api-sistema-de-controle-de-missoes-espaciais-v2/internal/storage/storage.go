package storage

import (
	"api-sistema-de-controle-de-missoes-espaciais-v2/internal/models"
	"slices"
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

func (st *Storage) GetAllShips() []*models.Ship {
	st.gate.RLock()
	defer st.gate.RUnlock()
	return st.ships
}

func (st *Storage) GetAllActives() []*models.Ship {
	st.gate.RLock()
	defer st.gate.RUnlock()
	var actives []*models.Ship

	for _, sh := range st.ships {
		if sh.Active {
			actives = append(actives, sh)
		}
	}
	return actives
}

func (st *Storage) find(name string) (*models.Ship, bool) {
	for _, sh := range st.ships {
		if sh.Name == name {
			return sh, true
		}
	}
	return nil, false
}

func (st *Storage) ExistShip(name string) (*models.Ship, bool) {
	st.gate.RLock()
	defer st.gate.RUnlock()
	sh, exist := st.find(name)

	return sh, exist
}

func (st *Storage) DeleteShip(name string) (*models.Ship, bool) {
	st.gate.Lock()
	defer st.gate.Unlock()
	sh, exist := st.find(name)
	if !exist {
		return sh, exist
	}

	st.ships = slices.DeleteFunc(st.ships, func(sh *models.Ship) bool {
		return sh.Name == name
	})

	return sh, exist
}

func (st *Storage) UpdateShip(name string, newData *models.Ship) (*models.Ship, bool) {

	st.gate.Lock()
	defer st.gate.Unlock()

	sh, exist := st.find(name)
	if !exist {
		return nil, false
	}

	if newData.Name != "" && newData.Name != sh.Name {
		sh.Name = newData.Name
	}

	if newData.Status != "" && newData.Status != sh.Status {
		sh.Status = newData.Status
	}

	sh.Active = newData.Active

	return sh, true
}
