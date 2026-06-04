package main

import (
	"fmt"
	"log"
	"net"
)

var ips []net.IP

func main() {
	fmt.Println("Função main")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func init() {
	fmt.Println("Função init")
	getIP, err := net.LookupIP("consultorfit.com.br")

	if err != nil {
		log.Fatal(err)
	}
	ips = getIP
}
