package servicesIP

import (
	"api-busca-ip-serve/internal/ip/models"
	"errors"
	"net"
)

type Service struct{}

func (s Service) GetIp(site string) (*models.ResponseIP, error) {
	ips, err := net.LookupIP(site)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	resIP := models.ResponseIP{
		HostName: site,
		ListIPs:  ips,
	}

	return &resIP, nil
}

func NewSerice() *Service {
	return &Service{}
}
