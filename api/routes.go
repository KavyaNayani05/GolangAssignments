package api

import (
	"database/sql"
	"net/http"
)

func RegisterRouter(db *sql.DB) {
	http.HandleFunc("/create", CreateHandler(db))
	//http.HandleFunc("/update", UpdateHandler(db))

}
