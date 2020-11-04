package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"udemy-go-books/db"
	"udemy-go-books/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var books []models.Book
	db := db.Init()

	rows, err := db.Query("select * from book")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.YearLaunch)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// param := mux.Vars(r)
	// id, err := strconv.Atoi(param["id"])
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	// var book models.Book
	// _ = json.NewDecoder(r.Body).Decode(&book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// var book models.Book
	// _ = json.NewDecoder(r.Body).Decode(&book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// param := mux.Vars(r)
	// id, err := strconv.Atoi(param["id"])
}
