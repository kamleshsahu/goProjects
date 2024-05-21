package main

type User struct {
	Id          int    `uri:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Age         int    `json:"age"`
	ContactNo   string `json:"contactNo"`
}

type CustomError interface {
}
