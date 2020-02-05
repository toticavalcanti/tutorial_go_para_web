package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
	"html/template"
 )	
//globals variables
var client *redis.Client
var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
var templates *template.Template

func main(){
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")	
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

 //request hello handle
 func indexGetHandler(w http.ResponseWriter, r *http.Request){
 	session, _ := store.Get(r, "session")
 	_, ok := session.Values["username"]
 	if !ok {
 		http.Redirect(w, r, "/login", 302)
 		return
 	}
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

func loginGetHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	hash, err := client.Get("user: " + username).Bytes()
	if err != nil{
		return
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil{
		return
	}
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return
	}
	client.Set("user: " + username, hash, 0)
	http.Redirect(w, r, "/login", 302)
}