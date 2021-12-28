package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Representa a URL para comunicacao com a API
	APIURL = ""
	//Onde a Aplicacao está rodando
	Porta = 0
	//Utilizado para autenticar o cookie
	HashKey []byte
	//Utilizado para criptografar os dados do cookie
	BlockKey []byte
)

//Inicializa as variaveis de ambiente
func Carregar() {
	var erro error

	if godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
