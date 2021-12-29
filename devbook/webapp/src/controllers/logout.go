package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

//Remover os dados de autenticacao do Usuario salvos no browser
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	status := 302

	cookies.Deletar(w)
	http.Redirect(w, r, "/login", status)
}
