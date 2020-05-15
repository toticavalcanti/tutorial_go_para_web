package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
	"html/template"
 )	
//globals variables
var client *redis.Client
var templates *template.Template

func main(){
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET") 
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8001", nil)
}

 //request hello handle
 func indexGetHandler(w http.ResponseWriter, r *http.Request){
 	comments, err := client.LRange("comments", 0, 10).Result()
 	if err != nil{
 		return
 	}
 	templates.ExecuteTemplate(w, "index.html", comments)
 }

  func indexPostHandler(w http.ResponseWriter, r *http.Request){
 	r.ParseForm()
 	comment := r.PostForm.Get("comment")
 	client.LPush("comments", comment)
 	http.Redirect(w, r, "/", 302)
 }

 //request contact page handle
func contactHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "contact.html", "This is the contact page!")
 }
 
 //request about page handle
 func aboutHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "about.html", "This is the about page!")
 }