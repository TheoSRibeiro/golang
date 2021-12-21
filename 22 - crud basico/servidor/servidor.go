package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//CriarUsuario insere usr no DB
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usr usuario

	if erro = json.Unmarshal(corpoRequisicao, &usr); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}
	fmt.Println(usr)

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao se conectar no banco de dados!"))
		return
	}
	defer db.Close()

	//Prepare Statement
	statement, erro := db.Prepare("insert into usuarios(nome, email) values (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usr.Nome, usr.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	//Status Codes
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

}

// BuscarUsuarios traz todos os usuários salvos no BD
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	//conexao com o BD
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao se conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	//select * from usuarios
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usr usuario

		if erro := linhas.Scan(&usr.ID, &usr.Nome, &usr.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usr)
	}
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON!"))
		return
	}

}

// BuscarUsuarios traz um usuario especifico salvo no BD
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
