package main

import(
	"fmt"
	"log"
	"net/http"
 )	

 //request handle
 func handler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "Hello world!")
 }

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}