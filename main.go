package main

import (
	"fmt"
	"log"
	"net/http"
	"udemy-go-books/routes"

	"github.com/gorilla/handlers"
)

func main() {
	route := routes.LoadRoute()

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(route)))
}
