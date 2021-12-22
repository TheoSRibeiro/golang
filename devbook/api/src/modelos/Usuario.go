package modelos

import "time"

// letra maiuscula para ser exportado
// Usuario representa um usr utilizando a rede social
type Usuario struct {
	ID          uint64    `json:"id,omitempty"`
	Nome        string    `json:"nome,omitempty"`
	Nick        string    `json:"nick,omitempty"`
	Email       string    `json:"email,omitempty"`
	Senha       string    `json:"senha,omitempty"`
	DataCriacao time.Time `json:"dataCriacao,omitempty"`
}
