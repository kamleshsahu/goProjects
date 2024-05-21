package models

type Patient struct {
	User
	Bookings map[int]int
}
