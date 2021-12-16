package app

import "github.com/urfave/cli"

//Gerar vai retornas a aplicacao de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de IPs e nomes de Servidor na internet"

	return app
}
