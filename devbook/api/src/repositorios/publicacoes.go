package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Representa um repositorio de publicacoes
type Publicacoes struct {
	db *sql.DB
}

//Criacao de um repositorio de publicacoes
func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Insere uma publicacao no DB
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(`
		insert into publicacoes(titulo, conteudo, autor_id) values(?, ?, ?)
	`)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

//Retorna uma publicacao do DB
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick
		from publicacoes p inner join usuarios u on u.id = p.autor_id
		where p.id = ?`, publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.DataCriacao,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

//Retorna as publicacoes dos usuarios seguidos e do proprio usuario que fez a requisicao
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick 
	from publicacoes p 
	inner join usuarios u on u.id = p.autor_id 
	inner join seguidores s on p.autor_id = s.usuario_id 
	where u.id = ? or s.seguidor_id = ?
	order by 1 desc`,
		usuarioID, usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.DataCriacao,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicaoId uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare(`
	update publicacoes 
	set titulo = ?, conteudo = ? 
	where id = ?`,
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicaoId); erro != nil {
		return erro
	}

	return nil
}
