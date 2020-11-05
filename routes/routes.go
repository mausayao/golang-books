package routes

import (
	"udemy-go-books/controllers"
	"udemy-go-books/db"

	"github.com/gorilla/mux"
)

var bookController = controllers.BookController{}
var bd = db.Init()

func LoadRoute() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/books", bookController.GetBooks(bd)).Methods("GET")
	router.HandleFunc("/books/{id}", bookController.GetBook(bd)).Methods("GET")
	router.HandleFunc("/books", bookController.AddBook(bd)).Methods("POST")
	router.HandleFunc("/books", bookController.UpdateBook(bd)).Methods("PUT")
	router.HandleFunc("/books/{id}", bookController.DeleteBook(bd)).Methods("DELETE")

	return router

}
