package dataservice

import (
	"database/sql"
	"encoding/json"
	"golang/model"
	"net/http"
)

func CreateCustomer(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var customer model.Customer
	//var x = 10
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {

		return err
	}
	x := "SELECT quantity FROM product WHERE Pname = ?"

	conditionValue := customer.ProductName
	var result int
	err2 := db.QueryRow(x, conditionValue).Scan(&result)
	if err2 != nil {
		return err2
	}

	if result >= customer.Quantity && customer.Quantity > 0 {

		query := "Insert INTo customer (Cname,phnum,email,Pname,quantity) VALUES (?,?,?,?,?) "

		_, err := db.Exec(query, customer.CustomerName, customer.PhoneNumber, customer.Email, customer.ProductName, customer.Quantity)
		if err != nil {
			return err
		}
		updateQuery := "UPDATE product SET quantity=? WHERE Pname=?"
		_, err1 := db.Exec(updateQuery, result-customer.Quantity, customer.ProductName)
		if err1 != nil {
			return err1
		}
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
	return nil
}
