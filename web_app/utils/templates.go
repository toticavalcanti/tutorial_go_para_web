package utils

import(
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates(pattern string){
	templates = template.Must(template.ParseGlob(pattern))
}

func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}){
	templates.ExecuteTemplate(w, tmpl, data)	
}