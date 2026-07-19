package models

import "net"

type RequestIP struct {
	HostName string `json:"hostName"`
}

type ResponseIP struct {
	HostName string   `json:"hostName"`
	ListIPs  []net.IP `json:"listIPs"`
}
