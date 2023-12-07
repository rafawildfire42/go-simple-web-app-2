package main

import (
	"net/http"
	"vuewebapp/routes"
)

func main() {
	routes.HandleRoutes()
	http.ListenAndServe(":8000", nil)
}
