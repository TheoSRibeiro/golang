package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//Insere os templates html na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

//renderiza uma pagina html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)
}
