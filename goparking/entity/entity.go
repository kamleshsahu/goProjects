package entity

type Car struct {
	CarNumber string
	Color     string
}

type Slot struct {
	Id        int
	IsBlocked bool
	Car       *Car
}
