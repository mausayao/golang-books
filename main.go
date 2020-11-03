package main

import (
	"log"
	"net/http"
	"udemy-go-books/controllers"
	"udemy-go-books/routes"
)

func main() {
	controllers.LoadData()
	route := routes.LoadRoute()
	log.Fatal(http.ListenAndServe(":8000", route))
}
