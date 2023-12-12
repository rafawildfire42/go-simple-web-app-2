package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"vuewebapp/models"
)

func PageAddSubjectView(w http.ResponseWriter, r *http.Request) {

	studentID := r.URL.Query().Get("studentID")

	data := struct {
		Action    string
		StudentID string
	}{
		Action:    "create",
		StudentID: studentID,
	}

	temp.ExecuteTemplate(w, "AddOrEditSubject", data)

}

func CreateOrEditSubjectView(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		title := r.FormValue("Title")
		score1Str := r.FormValue("Score1")
		score2Str := r.FormValue("Score2")
		score3Str := r.FormValue("Score3")
		score4Str := r.FormValue("Score4")

		score1, err := strconv.ParseFloat(score1Str, 64)
		if err != nil {
			fmt.Println("Erro durante a conversão do score1:", err)
		}

		score2, err := strconv.ParseFloat(score2Str, 64)
		if err != nil {
			fmt.Println("Erro durante a conversão do score2:", err)
		}

		score3, err := strconv.ParseFloat(score3Str, 64)
		if err != nil {
			fmt.Println("Erro durante a conversão do score3:", err)
		}

		score4, err := strconv.ParseFloat(score4Str, 64)
		if err != nil {
			fmt.Println("Erro durante a conversão do score4:", err)
		}

		studentID := r.FormValue("StudentID")
		subjectID := r.URL.Query().Get("subjectID")

		fmt.Println(studentID)
		fmt.Println(subjectID)

		if subjectID != "" {
			models.EditSubject(title, subjectID, score1, score2, score3, score4)
		} else {
			models.CreateSubject(title, studentID, score1, score2, score3, score4)
		}

		http.Redirect(w, r, "/", 301)
	}
}
