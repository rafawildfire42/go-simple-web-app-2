package models

import (
	"database/sql"
	"fmt"
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
}

type Discipline struct {
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
		firstName, lastName, serieType, email string
		id, serie, age                        int
	)

	query := fmt.Sprintf("SELECT * FROM students WHERE id=%s", studentID)

	studentRows, err := dbData.Query(query)

	if err != nil {
		panic(err.Error())
	}

	if studentRows.Next() {
		studentRows.Scan(&id, &firstName, &lastName, &serie, &serieType, &age, &email)

		student.ID = id
		student.FirstName = firstName
		student.LastName = lastName
		student.Serie = serie
		student.SerieType = serieType
		student.Age = age
		student.Email = email

	} else {
		fmt.Println("Nenhum estudante encontrado para o ID 4.")
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
			firstName, lastName, serieType, email string
			id, serie, age                        int
		)

		err := allStudents.Scan(&id, &firstName, &lastName, &serie, &serieType, &age, &email)

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

		students = append(students, student)

	}

	return students
}

func CreateStudent(firstName, lastName, serieType, email string, serie, age int) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	var createStudent *sql.Stmt
	var err error

	createStudent, err = dbData.Prepare("INSERT INTO students (FirstName, LastName, Serie, SerieType, Age, email) VALUES (?, ?, ?, ?, ?, ?);")

	if err != nil {
		panic(err.Error())
	}

	_, err = createStudent.Exec(firstName, lastName, serie, serieType, age, email)

	if err != nil {
		panic(err.Error())
	}

}

func EditStudent(firstName, lastName, serieType, email, studentID string, serie, age int) {

	dbData := db.ConnectDatabase()
	defer dbData.Close()

	var editStudent *sql.Stmt
	var err error

	editStudent, err = dbData.Prepare("UPDATE students SET FirstName = ?, LastName = ?, Serie = ?, SerieType = ?, Age = ?, email = ? WHERE ID = ?;")

	if err != nil {
		panic(err.Error())
	}

	_, err = editStudent.Exec(firstName, lastName, serie, serieType, age, email, studentID)

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

func GetDisciplines(studentID int) []Discipline {

	discipline := []Discipline{
		{ID: 1, Title: "Matemática", Score1: 9, Score2: 10, Score3: 10, Score4: 9, FinalMean: 9.5, StudentID: studentID},
	}

	return discipline

}
