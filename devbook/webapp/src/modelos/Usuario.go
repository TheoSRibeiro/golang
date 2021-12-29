package modelos

import "time"

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	DataCriacao time.Time    `json:"dataCriacao"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	publicacoes []Publicacao `json:"publicacoes"`
}
