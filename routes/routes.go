package routes

import (
	"net/http"
	"vuewebapp/controllers"
)

func HandleRoutes() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", controllers.IndexView)
	http.HandleFunc("/student", controllers.StudentView)
	http.HandleFunc("/add-student", controllers.PageCreateStudentView)
	http.HandleFunc("/edit-student", controllers.PageEditStudentView)
	http.HandleFunc("/createOrEdit", controllers.CreateOrEditStudentView)
	// http.HandleFunc("/edit", controllers.EditStudentView)
}
