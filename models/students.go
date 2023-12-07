package models

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Serie     int
	SerieType string
	Age       int
	email     string
}

type Test struct {
	ID         int
	Discipline string
	Score      float64
	Student    Student
	Period     int
}

func GetStudents() []Student {

	students := []Student{
		{ID: 1, FirstName: "Rafael", LastName: "Fontenele", Serie: 4, SerieType: "Médio"},
		{ID: 2, FirstName: "Maria", LastName: "Cecília", Serie: 3, SerieType: "Fundamental"},
	}

	return students
}
