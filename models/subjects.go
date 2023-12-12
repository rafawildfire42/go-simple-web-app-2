package models

import (
	"fmt"
	"log"
	"strconv"
	"vuewebapp/db"
)

type Subject struct {
	ID        int
	Title     string
	Score1    float64
	Score2    float64
	Score3    float64
	Score4    float64
	StudentID int
}

func GetDisciplineByStudentID(studentID string) Subject {
	var id int
	var title string
	var score1, score2, score3, score4 float64
	var subject Subject

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	query := fmt.Sprintf("SELECT * FROM subjects WHERE StudentID=%s", studentID)

	fmt.Println(query)

	rows, err := dbData.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id, &title, &score1, &score2, &score3, &score4, &studentID)
		if err != nil {
			log.Fatal(err)
		}

		studentIDInt, err := strconv.Atoi(studentID)
		if err != nil {
			log.Fatal(err)
		}

		subject.ID = id
		subject.Title = title
		subject.Score1 = score1
		subject.Score2 = score2
		subject.Score3 = score3
		subject.Score4 = score4
		subject.StudentID = studentIDInt
	}

	return subject
}

func CreateSubject(title, studentID string, score1, score2, score3, score4 float64) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	createSubject, err := dbData.Prepare("INSERT INTO subjects (Title, Score1, Score2, Score3, Score4, StudentID) VALUES (?, ?, ?, ?, ?, ?);")

	if err != nil {
		panic(err.Error())
	}

	_, err = createSubject.Exec(title, score1, score2, score3, score4, studentID)

	if err != nil {
		panic(err.Error())
	}

}

func EditSubject(title, subjectID string, score1, score2, score3, score4 float64) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	editSubject, err := dbData.Prepare("UPDATE subjects SET Title = ?, Score1 = ?, Score2 = ?, Score3 = ?, Score4 = ? WHERE ID = ?;")

	if err != nil {
		panic(err.Error())
	}

	_, err = editSubject.Exec(title, score1, score2, score3, score4, subjectID)

	if err != nil {
		panic(err.Error())
	}

}
