package interfaces

import (
	"awesomeProject/goparking/entity"
)

type IParkingService interface {
	AllotParkingSpace(carNumber string, color string) int
	ReleaseParkingSpace(carNumber string)
	ReleaseParkingSpaceBySlotId(slotId int)
	GetAllCarsByColor(color string) []entity.Car
}

type ISlotService interface {
	GetNextSlot() (int, bool)
	BlockSlot(car entity.Car, slotId int) bool
	ReleaseSlot(slotId int) bool
	GetCarSlot(carNumber string) (int, bool)
	GetAllCarsByColor(color string) []entity.Car
}
