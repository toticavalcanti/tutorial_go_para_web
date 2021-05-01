package main

import (
	"net/http"

	"example.com/web-app/models"
	"example.com/web-app/routes"
	"example.com/web-app/utils"
)

func main() {
	models.Init()
	utils.LoadTemplates("templates/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8001", nil)
}
