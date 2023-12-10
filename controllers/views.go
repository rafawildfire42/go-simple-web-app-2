package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"vuewebapp/models"
)

type Data struct {
	Operation string
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func IndexView(w http.ResponseWriter, r *http.Request) {

	students := models.GetStudents()

	temp.ExecuteTemplate(w, "Index", students)

}

func StudentView(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	discplines := models.GetDisciplines(id)

	temp.ExecuteTemplate(w, "Student", discplines)

}

func EditStudentView(w http.ResponseWriter, r *http.Request) {

	data := Data{
		Operation: "edit",
	}

	temp.ExecuteTemplate(w, "AddOrEditStudent", data)

}

func AddStudentView(w http.ResponseWriter, r *http.Request) {

	data := Data{
		Operation: "add",
	}

	temp.ExecuteTemplate(w, "AddOrEditStudent", data)

}

func CreateStudentView(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		serie := r.FormValue("serie")
		serieType := r.FormValue("serieType")
		age := r.FormValue("age")
		email := r.FormValue("email")

		ageConverted, errAge := strconv.Atoi(age)
		serieConverted, errSerie := strconv.Atoi(serie)

		if errAge != nil || errSerie != nil {
			fmt.Println("Erro durante a conversão.")
		}

		models.CreateStudent(firstName, lastName, serieType, email, serieConverted, ageConverted)

	}

	http.Redirect(w, r, "/", 301)

}
