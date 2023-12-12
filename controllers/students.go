package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"vuewebapp/models"
)

func IndexView(w http.ResponseWriter, r *http.Request) {

	Students := models.ListStudents()

	temp.ExecuteTemplate(w, "Index", Students)

}

func StudentView(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("id")

	subject := models.GetDisciplineByStudentID(studentID)
	student := models.GetStudent(studentID)

	data := struct {
		Subject models.Subject
		Student models.Student
	}{
		Subject: subject,
		Student: student,
	}

	temp.ExecuteTemplate(w, "Student", data)
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

func DeleteStudentView(w http.ResponseWriter, r *http.Request) {

	studentID := r.URL.Query().Get("id")

	models.DeleteStudent(studentID)

	http.Redirect(w, r, "/", 301)

}

func CreateOrEditStudentView(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")
		serie := r.FormValue("serie")
		serieType := r.FormValue("serieType")
		age := r.FormValue("age")
		email := r.FormValue("email")
		gender := r.FormValue("gender")

		ageConverted, errAge := strconv.Atoi(age)
		serieConverted, errSerie := strconv.Atoi(serie)

		if errAge != nil || errSerie != nil {
			fmt.Println("Erro durante a conversão.")
		}

		studentID := r.FormValue("ID")

		if studentID != "0" {
			models.EditStudent(firstName, lastName, serieType, email, studentID, gender, serieConverted, ageConverted)
		} else {
			models.CreateStudent(firstName, lastName, serieType, email, gender, serieConverted, ageConverted)
		}

	}

	http.Redirect(w, r, "/", 301)

}
