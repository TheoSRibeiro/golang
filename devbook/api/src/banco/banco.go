package banco

import (
	"api/src/config"
	"database/sql"
)

// Abrir a conexao com o DB e a retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
