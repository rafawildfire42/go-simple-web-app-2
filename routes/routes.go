package routes

import (
	"net/http"
	"vuewebapp/controllers"
)

func HandleRoutes() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", controllers.IndexView)
	http.HandleFunc("/student", controllers.StudentView)
}
