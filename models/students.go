package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"vuewebapp/db"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Serie     int
	SerieType string
	Age       int
	Email     string
	Gender    string
}

type Subject struct {
	ID        int
	Title     string
	Score1    float64
	Score2    float64
	Score3    float64
	Score4    float64
	FinalMean float64
	StudentID int
}

func GetStudent(studentID string) Student {
	dbData := db.ConnectDatabase()
	defer dbData.Close()

	var student Student

	var (
		firstName, lastName, serieType, email, gender string
		id, serie, age                                int
	)

	query := fmt.Sprintf("SELECT * FROM students WHERE id=%s", studentID)

	studentRows, err := dbData.Query(query)

	if err != nil {
		panic(err.Error())
	}

	if studentRows.Next() {
		studentRows.Scan(&id, &firstName, &lastName, &serie, &serieType, &age, &email, &gender)

		student.ID = id
		student.FirstName = firstName
		student.LastName = lastName
		student.Serie = serie
		student.SerieType = serieType
		student.Age = age
		student.Email = email
		student.Gender = gender

	} else {
		fmt.Println("Nenhum estudante encontrado.")
	}

	return student

}

func ListStudents() []Student {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	allStudents, err := dbData.Query("select * from students")

	if err != nil {
		panic(err.Error())
	}

	var student Student
	var students []Student

	for allStudents.Next() {
		var (
			firstName, lastName, serieType, email, gender string
			id, serie, age                                int
		)

		err := allStudents.Scan(&id, &firstName, &lastName, &serie, &serieType, &age, &email, &gender)

		if err != nil {
			panic(err.Error())
		}

		student.ID = id
		student.FirstName = firstName
		student.LastName = lastName
		student.Serie = serie
		student.SerieType = serieType
		student.Age = age
		student.Email = email
		student.Gender = gender

		students = append(students, student)

	}

	return students
}

func CreateStudent(firstName, lastName, serieType, email, gender string, serie, age int) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	var createStudent *sql.Stmt
	var err error

	createStudent, err = dbData.Prepare("INSERT INTO students (FirstName, LastName, Serie, SerieType, Age, Email, Gender) VALUES (?, ?, ?, ?, ?, ?);")

	if err != nil {
		panic(err.Error())
	}

	_, err = createStudent.Exec(firstName, lastName, serie, serieType, age, email)

	if err != nil {
		panic(err.Error())
	}

}

func EditStudent(firstName, lastName, serieType, email, studentID, gender string, serie, age int) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	var editStudent *sql.Stmt
	var err error

	editStudent, err = dbData.Prepare("UPDATE students SET FirstName = ?, LastName = ?, Serie = ?, SerieType = ?, Age = ?, Email = ?, Gender = ? WHERE ID = ?;")

	if err != nil {
		panic(err.Error())
	}

	_, err = editStudent.Exec(firstName, lastName, serie, serieType, age, email, gender, studentID)

	if err != nil {
		panic(err.Error())
	}

}

func DeleteStudent(studentID string) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	deleteStudent, err := dbData.Prepare("DELETE FROM students WHERE id = ?;")

	if err != nil {
		panic(err.Error())
	}

	_, err = deleteStudent.Exec(studentID)

	if err != nil {
		panic(err.Error())
	}

}

func GetDiscipline(studentID string) Subject {

	var id int
	var score1, score2, score3, score4 float64
	var subject Subject

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	query := fmt.Sprintf("SELECT * FROM subjects WHERE StudentID=%s", studentID)

	dbData.Query(query)

	studentRows, _ := dbData.Query(query)

	for studentRows.Next() {
		studentID, _ := strconv.Atoi(studentID)
		studentRows.Scan(&id, &score1, &score2, &score3, &score4, &studentID)

		subject.ID = id
		subject.Score1 = score1
		subject.Score2 = score1
		subject.Score3 = score1
		subject.Score4 = score1
		subject.StudentID = studentID

	}

	return subject

}
