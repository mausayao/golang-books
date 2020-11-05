package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"udemy-go-books/models"
	bookrepository "udemy-go-books/repository/book"
	"udemy-go-books/utils"

	"github.com/gorilla/mux"
)

type BookController struct{}

func (c BookController) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		books := []models.Book{}

		bookRepo := bookrepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "Application/json")
		utils.SendSuccess(w, books)
	}
}

func (c BookController) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		var book models.Book
		var error models.Error
		id, err := strconv.Atoi(param["id"])

		bookRepo := bookrepository.BookRepository{}
		book, err = bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}

		}

		w.Header().Set("Content-Type", "Application/json")
		utils.SendSuccess(w, book)

	}
}

func (c BookController) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.YearLaunch == "" {
			error.Message = "Enter missing fields"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "Application/json")
		utils.SendSuccess(w, bookID)

	}
}

func (c BookController) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.YearLaunch == "" {
			error.Message = "All fields are required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookrepository.BookRepository{}
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "Application/json")
		utils.SendSuccess(w, rowsUpdated)

	}
}

func (c BookController) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		param := mux.Vars(r)
		id, _ := strconv.Atoi(param["id"])

		bookRepo := bookrepository.BookRepository{}
		rowsDeleted, err := bookRepo.DeleteBook(db, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "Application/json")
		utils.SendSuccess(w, rowsDeleted)
	}
}
