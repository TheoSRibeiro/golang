package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Autenticacao do usuario na API
func Login(w http.ResponseWriter, r *http.Request) {

	//ler corpo da requisicao
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usr modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usr); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	usuarioSalvoNoDB, erro := repositorio.BuscarPorEmail(usr.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoDB.Senha, usr.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, _ := autenticacao.CriarToken(usuarioSalvoNoDB.ID)
	fmt.Println(token)
	w.Write([]byte(token))
}
