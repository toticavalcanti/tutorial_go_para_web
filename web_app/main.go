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
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

 //request hello handle
 func hiHandler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "Hi everyone!")
 }

 //request goodbye handle
 func byeHandler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "bye bye!")
 }