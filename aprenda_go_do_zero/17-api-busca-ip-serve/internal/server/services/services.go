package servicesServer

import (
	"api-busca-ip-serve/internal/server/models"
	"errors"
	"net"
)

type Service struct{}

func (s Service) GetServer(site string) (*models.ResponseServer, error) {
	servers, err := net.LookupNS(site)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	resServer := models.ResponseServer{
		HostName:    site,
		ListServers: servers,
	}

	return &resServer, nil
}

func NewSerice() *Service {
	return &Service{}
}
