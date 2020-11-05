package bookrepository

import (
	"database/sql"
	"udemy-go-books/models"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select * from book")

	if err != nil {
		return []models.Book{}, err
	}

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.YearLaunch)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil

}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {

	row := db.QueryRow("select * from book where id=$1", id)
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.YearLaunch)

	return book, err
}

func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {

	err := db.QueryRow(
		"insert into book (title, author, year_launch) values ($1, $2, $3) RETURNING id;", book.Title, book.Author, book.YearLaunch).Scan(&book.ID)

	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {

	result, err := db.Exec("update book set title=$1, author=$2, year_launch=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.YearLaunch, &book.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

func (b BookRepository) DeleteBook(db *sql.DB, id int) (int64, error) {

	result, err := db.Exec("delete from book where id=$1", id)

	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil

}
