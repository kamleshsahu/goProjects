package models

type User struct {
	Id      int
	Name    string
	Age     int
	Number  []string
	Address Address
}

type Address struct {
	AddressLine1 string
	AddressLine2 string
	City         string
	State        string
	Country      string
	Location     Location
}

type Location struct {
	Lat, Long float32
}
