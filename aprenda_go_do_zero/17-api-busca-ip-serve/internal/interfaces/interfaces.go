package interfaces

import "api-busca-ip-serve/internal/ip/models"

type Services interface {
	GetIp(site string) (*models.ResponseIP, error)
}
