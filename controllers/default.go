package controllers

import (
	"text/template"
	"vuewebapp/models"
)

type Data struct {
	Action  string
	Student models.Student
}

var temp = template.Must(template.New("").Funcs(template.FuncMap{
	"add": add,
	"div": div,
}).ParseGlob("templates/*.html"))

func div(a, b float64) float64 {
	return a / b
}

func add(a, b, c, d float64) float64 {
	return a + b + c + d
}
