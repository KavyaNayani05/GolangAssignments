package api

import (
	"database/sql"
	"golang/dataservice"
	"net/http"
)

/*func CreateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateBook(db, w, r)

}

func UpdateBookLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.UpdateBook(db, w, r)

}*/

func CreateCustomerLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.CreateCustomer(db, w, r)

}
