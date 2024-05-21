package entity

type Floor struct {
	Id int
}

type Building struct {
	Floors []Floor
}

type InternalButton struct {
	AvailableButtons []int
}

type Direction int

const (
	UP = iota
	DOWN
)

type ElevatorState int

const (
	STILL = iota
	MOVING
)

type ElevatorCar struct {
	Id            int
	CurrentFloor  int
	Direction     Direction
	ElevatorState ElevatorState
}
