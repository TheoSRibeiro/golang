package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?,?,?,?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

//Traz todos os usuarios que atendem o filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {

	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, dataCriacao from usuarios where nome LIKE ? or nick LIKE ? ",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.DataCriacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}
