package main

import (
	"fmt"
	"linha_comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	//Inicializar a aplicacao - fazer os comandos sejam reconhecidos pela linha de comando
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
