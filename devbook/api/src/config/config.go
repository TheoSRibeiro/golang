package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Conexao com o MySQL
	ConexaoBanco = ""

	// Porta onde a API vai estar rodando
	Porta = 0
)

// Inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	//carregar os dados do .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// converter a porta de string para int
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	//se a porta 5000 estiver ocupada, vai carregar a API na porta 9000
	if erro != nil {
		Porta = 9000
	}

	ConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
