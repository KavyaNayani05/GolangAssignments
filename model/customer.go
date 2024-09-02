package model

type Customer struct {
	CustomerName string `json:"Cname"`
	PhoneNumber  int64  `json:"phnum"`
	Email        string `json:"email"`
	ProductName  string `json:"Pname"`
	Quantity     int    `json:"quantity"`
}
type Product struct {
	ProductID   int     `json:"PID"`
	ProductName string  `json:"Pname"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
