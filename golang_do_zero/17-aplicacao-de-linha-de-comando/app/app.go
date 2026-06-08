package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Faz alguma coisa"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca os ips da urls digitadas",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "consultorfit.com.br",
				},
			},
			Action: buscarIPs,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		println(ip.String())
	}
}
