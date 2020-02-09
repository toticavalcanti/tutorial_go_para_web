package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
 )	

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/hi", hiHandler).Methods("GET")
	r.HandleFunc("/bye", byeHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

 //request index page handle
 func indexHandler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "This is the index page of your web aplication!")
 }

 //request hi everyone! page handle
 func hiHandler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "Hi everyone!")
 }

 //request bye bye! page handle
 func byeHandler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "bye bye!")
 }