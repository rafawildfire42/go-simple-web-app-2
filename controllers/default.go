package controllers

import (
	"text/template"
	"vuewebapp/models"
)

type Data struct {
	Action  string
	Student models.Student
}

var temp = template.Must(template.ParseGlob("templates/*.html"))
