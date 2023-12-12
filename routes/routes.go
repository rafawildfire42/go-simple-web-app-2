package routes

import (
	"net/http"
	"vuewebapp/controllers"
)

func HandleRoutes() {
	// Carregar os arquivos estáticos
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Página inicial
	http.HandleFunc("/", controllers.IndexView)

	// Visualizar dados do aluno
	http.HandleFunc("/student", controllers.StudentView)

	// Formulários para adicionar ou editar aluno e disciplina
	http.HandleFunc("/add-student", controllers.PageCreateStudentView)
	http.HandleFunc("/edit-student", controllers.PageEditStudentView)
	http.HandleFunc("/add-subject", controllers.PageAddSubjectView)

	// Apagar, deletar ou editar aluno
	http.HandleFunc("/createOrEdit", controllers.CreateOrEditStudentView)
	http.HandleFunc("/delete", controllers.DeleteStudentView)

	// Apagar, deletar ou editar disciplina

}
