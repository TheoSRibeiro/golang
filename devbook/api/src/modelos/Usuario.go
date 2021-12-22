package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

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

// Chamar os metodos para validar e formatar o usuario recebido no Controller
func (usr *Usuario) Preparar(etapa string) error {
	if erro := usr.validar(etapa); erro != nil {
		return erro
	}

	usr.formatar()
	return nil
}

func (usr *Usuario) validar(etapa string) error {
	if usr.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco!")
	}

	if usr.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco!")
	}

	if usr.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco!")
	}

	if erro := checkmail.ValidateFormat(usr.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido!")
	}

	if etapa == "cadastro" && usr.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco!")
	}

	return nil
}

func (usr *Usuario) formatar() {
	usr.Nome = strings.TrimSpace(usr.Nome)
	usr.Nick = strings.TrimSpace(usr.Nick)
	usr.Email = strings.TrimSpace(usr.Email)
}
