package modelos

import "time"

//Representa uma Publicacao feita por um usuario
type Publicacao struct {
	ID          uint64    `json:"id,omitempty"`
	Titulo      string    `json:"titulo,omitempty"`
	Conteudo    string    `json:"conteudo,omitempty"`
	AutorID     uint64    `json:"autorId,omitempty"`
	AutorNick   string    `json:"autorNick,omitempty"`
	Curtidas    uint64    `json:"curtidas"`
	DataCriacao time.Time `json:"dataCriacao,omitempty"`
}
