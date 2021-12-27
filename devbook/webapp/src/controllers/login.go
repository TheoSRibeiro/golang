package controllers

import "net/http"

//Carregar a tela de login
func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("tela de login"))
}
