package dataservice

import (
	"database/sql"
	"encoding/json"
	"golang/model"
	"net/http"
)

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {

		return err
	}

	query := "Insert INTo book (title,author,year) VALUES (?,?,?) "

	_, err := db.Exec(query, book.Title, book.Author, book.Year)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}

func UpdateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book model.Book
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&book); err != nil {

		return err
	}

	updateQuery := "UPDATE book SET author=?, year=? WHERE title=?"
	_, err1 := db.Exec(updateQuery, book.Author, book.Year, book.Title)
	if err1 != nil {
		return err1
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}
