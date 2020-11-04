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
	bd := db.Init()
	log.Println(bd.Ping())
	log.Fatal(http.ListenAndServe(":8000", route))
}
