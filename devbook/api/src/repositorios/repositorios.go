package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Repositorio de usuarios
type usuarios struct {
	db *sql.DB
}

// criacao de um repositorio de usuarios
func NovoRepositorioUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//METODOS

//Inserir um usuario no DB
func (u usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
