package entity

import "time"

type SeatCategory int

const (
	SILVER = iota
	PLATINUM
	GOLD
)

type BookingStatus int

const (
	PENDING = iota
	COMPLETED
	FAILED
)

type Seat struct {
	Id           int
	Row          int
	SeatCategory SeatCategory
}

type Screen struct {
	Id    int
	Seats []Seat
}

type Movie struct {
	Id   int
	Name string
}

type Show struct {
	Id        int
	ScreenId  int
	MovieId   int
	StartTime time.Time
	EndTime   time.Time
}

type Theatre struct {
	Id      int
	Screens []Screen
	Shows   []Show
	City    string
}

type Booking struct {
	Id        int
	ShowId    int
	Status    BookingStatus
	Seats     []Seat
	TheatreId int
}
