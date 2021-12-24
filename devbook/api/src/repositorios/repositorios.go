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

// Retornar um usuario do DB
func (repositorio usuarios) BuscarID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, dataCriacao from usuarios where id = ?", ID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.DataCriacao,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

// Atualizar os dados do usuario no DB
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {

	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ? ",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Excluir as informacoes dos usuarios no DB
func (repositorio usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

//Busca um usuario no DB e retorna o seu ID e senha com Hash
func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Permitir um usuario seguir outro
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores(usuario_id, seguidor_id) values(?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//Permitir um usr pare de seguir o outro
func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	statement.Close()

	if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

//Retorna todos os seguidores de 1 Usr
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.dataCriacao
		from usuarios u inner join seguidores s on u.id = s.seguidor_id
		where s.usuario_id = ?`,
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usr modelos.Usuario

		if erro = linhas.Scan(
			&usr.ID,
			&usr.Nome,
			&usr.Nick,
			&usr.Email,
			&usr.DataCriacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usr)
	}

	return usuarios, nil
}

//Retorna todos os seguidores que um determinado usuario esta seguindo
func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.dataCriacao
		from usuarios u inner join seguidores s on u.id = s.usuario_id 
		where s.seguidor_id = ?`, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usr modelos.Usuario

		if erro = linhas.Scan(
			&usr.ID,
			&usr.Nome,
			&usr.Nick,
			&usr.Email,
			&usr.DataCriacao,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usr)
	}

	return usuarios, nil
}
