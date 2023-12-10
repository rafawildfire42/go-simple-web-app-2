package routes

import (
	"net/http"
	"vuewebapp/controllers"
)

func HandleRoutes() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", controllers.IndexView)
	http.HandleFunc("/student", controllers.StudentView)
	http.HandleFunc("/add-student", controllers.AddStudentView)
	http.HandleFunc("/create", controllers.CreateStudentView)
	http.HandleFunc("/edit-student", controllers.EditStudentView)
}
