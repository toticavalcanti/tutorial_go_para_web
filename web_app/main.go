package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
 )	

var templates *template.template

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/goodbye", goodbyeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

 //request hello handle
 func indexHandler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "Hello world!")
 }