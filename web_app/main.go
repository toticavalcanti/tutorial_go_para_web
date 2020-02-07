package main

import(
	"net/http"
	"./routes"
	"./models"
	"./sessions"
	"./utils"
 )	

func main(){
	models.Init()
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
 	session, _ := sessions.Store.Get(r, "session")
 	_, ok := session.Values["username"]
 	if !ok {
 		http.Redirect(w, r, "/login", 302)
 		return
 	}
	handler.ServeHTTP(w, r)
	}
}

 func indexGetHandler(w http.ResponseWriter, r *http.Request) {
 	comments, err := models.GetComments()
 	if err != nil{
 		w.WriteHeader(http.StatusInternalServerError)
 		w.Write([]byte("Internal server error"))
 		return
 	}
 	utils.ExecuteTemplate(w, "index.html", comments)
 }

  func indexPostHandler(w http.ResponseWriter, r *http.Request){
 	r.ParseForm()
 	comment := r.PostForm.Get("comment")
 	err := models.PostComment(comment)
 	if err != nil{
 		w.WriteHeader(http.StatusInternalServerError)
 		w.Write([]byte("Internal server error"))
 		return
 	}
 	http.Redirect(w, r, "/", 302)
 }

func loginGetHandler(w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.AuthenticateUser(username, password)
	if err != nil{
		switch err {
		case models.ErrUserNotFound:
			utils.ExecuteTemplate(w, "login.html", "unknown user")
		case models.ErrInvalidLogin:
			utils.ExecuteTemplate(w, "login.html", "invalid login")
		default:
			w.WriteHeader(http.StatusInternalServerError)
 			w.Write([]byte("Internal server error"))
		}
		return
	}
	session, _ := sessions.Store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}

func registerGetHandler(w http.ResponseWriter, r *http.Request){
	utils.ExecuteTemplate(w, "register.html", nil)
}

func registerPostHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	err := models.RegisterUser(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
 		w.Write([]byte("Internal server error"))
 		return
	}
	http.Redirect(w, r, "/login", 302)
}