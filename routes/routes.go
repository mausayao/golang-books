package routes

import (
	"udemy-go-books/controllers"

	"github.com/gorilla/mux"
)

func LoadRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.AddBook).Methods("POST")
	router.HandleFunc("/books", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")

	return router

}
