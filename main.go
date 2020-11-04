package main

import (
	"log"
	"net/http"
	"udemy-go-books/controllers"
	"udemy-go-books/db"
	"udemy-go-books/routes"
)

func main() {
	controllers.LoadData()
	route := routes.LoadRoute()
	db.Init()
	log.Fatal(http.ListenAndServe(":8000", route))
}
