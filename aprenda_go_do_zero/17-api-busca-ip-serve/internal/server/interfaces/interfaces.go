package interfaces

import "api-busca-ip-serve/internal/server/models"

type Services interface {
	GetServer(site string) (*models.ResponseServer, error)
}
