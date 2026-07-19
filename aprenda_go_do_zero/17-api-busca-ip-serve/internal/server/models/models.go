package models

import "net"

type RequestServer struct {
	HostName string `json:"hostName"`
}

type ResponseServer struct {
	HostName    string    `json:"hostName"`
	ListServers []*net.NS `json:"listServers"`
}
