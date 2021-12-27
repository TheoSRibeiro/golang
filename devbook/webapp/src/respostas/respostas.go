package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

//Representa a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

//Retorna uma resposta em formato JSON para a requisicao
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

//Trata as requisicoes com status code 400 ou superiores
func TratarStatusCodeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
