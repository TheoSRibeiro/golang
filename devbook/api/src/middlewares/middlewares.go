package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

//verifica se o usr ao fazer a requisicao esta autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Autenticando...")
		next(w, r) //se autenticar, faz a funcao da rota
	}

}

//Escreve informacoes da requisicao no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
