package modelos

import (
	"net/http"
	"time"
)

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	DataCriacao time.Time    `json:"dataCriacao"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

//Faz 4 Requisicoes na API para montar o Usuario
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

}

func BuscarDadosUsuario(canal <-chan Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguidores(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarSeguindo(canal <-chan []Usuario, usuarioID uint64, r *http.Request) {

}

func BuscarPublicacoes(canal <-chan []Publicacao, usuarioID uint64, r *http.Request) {

}
