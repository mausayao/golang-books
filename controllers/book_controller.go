package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"udemy-go-books/db"
	"udemy-go-books/models"

	"github.com/gorilla/mux"
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

	rows.Close()

	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])

	var book models.Book

	if err != nil {
		log.Panic(err)
	}

	db := db.Init()

	row := db.QueryRow("select * from book where id=$1", id)

	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.YearLaunch)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(book)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)
	db := db.Init()

	err := db.QueryRow(
		"insert into book (title, author, year_launch) values ($1, $2, $3) RETURNING id;", book.Title, book.Author, book.YearLaunch).Scan(&bookID)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(bookID)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	db := db.Init()

	result, err := db.Exec("update book set title=$1, author=$2, year_launch=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.YearLaunch, &book.ID)

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(rowsUpdated)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// param := mux.Vars(r)
	// id, err := strconv.Atoi(param["id"])
}
