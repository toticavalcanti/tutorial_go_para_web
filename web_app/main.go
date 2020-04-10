package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
 )

var templates *template.Template

func main(){
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

 //request index page handle
 func indexHandler(w http.ResponseWriter, r *http.Request){
 	templates.ExecuteTemplate(w, "index.html", "This is the index page!")
 }

 //request contact page handle
 func contactHandler(w http.ResponseWriter, r *http.Request){
    templates.ExecuteTemplate(w, "contact.html", "This is the contact page!")
 }

 //request about page handle
 func aboutHandler(w http.ResponseWriter, r *http.Request){
    templates.ExecuteTemplate(w, "about.html", "This is the about page!")
 }