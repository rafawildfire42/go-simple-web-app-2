package controllers

import (
	"net/http"
	"text/template"
	"vuewebapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func IndexView(w http.ResponseWriter, r *http.Request) {

	students := models.GetStudents()

	temp.ExecuteTemplate(w, "Index", students)
}

func StudentView(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Student", nil)
}
