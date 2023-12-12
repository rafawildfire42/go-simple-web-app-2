package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"vuewebapp/models"
)

type Data struct {
	Action  string
	Student models.Student
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func IndexView(w http.ResponseWriter, r *http.Request) {

	students := models.ListStudents()

	temp.ExecuteTemplate(w, "Index", students)

}

func StudentView(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	discplines := models.GetDisciplines(id)

	temp.ExecuteTemplate(w, "Student", discplines)

}

func PageEditStudentView(w http.ResponseWriter, r *http.Request) {

	studentID := r.URL.Query().Get("id")

	student := models.GetStudent(studentID)

	data := Data{
		Action:  "edit",
		Student: student,
	}

	temp.ExecuteTemplate(w, "AddOrEditStudent", data)

}

func PageCreateStudentView(w http.ResponseWriter, r *http.Request) {

	student := models.Student{}

	data := Data{
		Action:  "create",
		Student: student,
	}

	temp.ExecuteTemplate(w, "AddOrEditStudent", data)

}

func CreateOrEditStudentView(w http.ResponseWriter, r *http.Request) {

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

		studentID := r.FormValue("ID")

		fmt.Println(studentID)

		if studentID != "" {
			models.EditStudent(firstName, lastName, serieType, email, studentID, serieConverted, ageConverted)
		} else {
			models.CreateStudent(firstName, lastName, serieType, email, serieConverted, ageConverted)
		}

	}

	http.Redirect(w, r, "/", 301)

}
