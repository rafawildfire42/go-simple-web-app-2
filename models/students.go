package models

import (
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

func GetStudents() []Student {

	// students := []Student{
	// 	{ID: 1, FirstName: "Rafael", LastName: "Fontenele", Serie: 4, SerieType: "Médio"},
	// 	{ID: 2, FirstName: "Maria", LastName: "Cecília", Serie: 3, SerieType: "Fundamental"},
	// }

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

	createStudent, err := dbData.Prepare("INSERT INTO students (FirstName, LastName, Serie, SerieType, Age, email) VALUES (?, ?, ?, ?, ?, ?);")

	if err != nil {
		panic(err.Error())
	}

	createStudent.Exec(firstName, lastName, serie, serieType, email, age)

}

func GetDisciplines(studentID int) []Discipline {

	discipline := []Discipline{
		{ID: 1, Title: "Matemática", Score1: 9, Score2: 10, Score3: 10, Score4: 9, FinalMean: 9.5, StudentID: studentID},
	}

	return discipline
}
