package models

type Ship struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Active bool   `json:"active"`
}
