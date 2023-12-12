package main

import (
	"log"
	"net/http"
	"vuewebapp/routes"

	"github.com/joho/godotenv"
	. "github.com/tbxark/g4vercel"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

func main() {
	routes.HandleRoutes()
	http.ListenAndServe(":8000", nil)
}
