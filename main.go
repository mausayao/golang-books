package main

import (
	"log"
	"net/http"
	"udemy-go-books/routes"
)

func main() {
	route := routes.LoadRoute()
	log.Fatal(http.ListenAndServe(":8000", route))
}
