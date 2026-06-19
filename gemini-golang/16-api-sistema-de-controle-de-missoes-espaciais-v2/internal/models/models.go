package models

import "sync"

type Ship struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Active bool   `json:"active"`
}

var gate sync.RWMutex

func (sh *Ship) UpName(name string) {
	gate.Lock()
	defer gate.Unlock()
	if name != sh.Name && name != "" {
		sh.Name = name
	}
}

func (sh *Ship) UpStatus(status string) {
	gate.Lock()
	defer gate.Unlock()
	if status != sh.Status && status != "" {
		sh.Status = status
	}
}

func (sh *Ship) UpActive(active bool) {
	gate.Lock()
	defer gate.Unlock()
	if active != sh.Active {
		sh.Active = active
	}
}
