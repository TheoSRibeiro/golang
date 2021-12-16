package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornas a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de IPs e nomes de Servidor na internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na Internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIPs,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
