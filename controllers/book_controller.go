package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"udemy-go-books/models"

	"github.com/gorilla/mux"
)

var books []models.Book

func LoadData() {
	books = append(books, models.Book{ID: 1, Title: "Golang pointers", Author: "Mr Golang", Year: "2010"},
		models.Book{ID: 2, Title: "Goroutines", Author: "Mr Goroutine", Year: "2011"},
		models.Book{ID: 3, Title: "Golang routers", Author: "Mr router", Year: "2012"},
		models.Book{ID: 4, Title: "Golang concurrency", Author: "Mr Currency", Year: "2013"},
		models.Book{ID: 5, Title: "Golang good parts", Author: "Mr Good", Year: "2014"})
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		panic(err.Error())
	}

	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
		}
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	for idx, item := range books {
		if item.ID == book.ID {
			books[idx] = book
		}
	}

	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete a books")
}
