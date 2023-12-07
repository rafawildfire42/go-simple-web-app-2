package routes

import (
	"net/http"
	"vuewebapp/controllers"
)

func HandleRoutes() {
	http.HandleFunc("/", controllers.IndexView)
	http.HandleFunc("/student", controllers.StudentView)
}
