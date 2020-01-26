package main

import(
	"fmt"
	"net/http"
 )	

 //request handle
 func handler(w http.ResponseWriter, r *http.Request){
 	fmt.Fprint(w, "Hello world!")
 }

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}