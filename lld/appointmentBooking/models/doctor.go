package models

type Doctor struct {
	User       User
	Rating     float32
	Speciality int
	SlotMap    map[int]int
}
