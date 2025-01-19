package service

import (
	"awesomeProject/goparking/entity"
	"awesomeProject/goparking/interfaces"
)

type ParkingService struct {
	slotService interfaces.ISlotService
}

func (p *ParkingService) GetAllCarsByColor(color string) []entity.Car {
	return p.slotService.GetAllCarsByColor(color)
}

func (p *ParkingService) ReleaseParkingSpaceBySlotId(slotId int) {
	p.slotService.ReleaseSlot(slotId)
}

func (p *ParkingService) AllotParkingSpace(carNumber string, color string) int {
	nextSlotId, hasNext := p.slotService.GetNextSlot()
	if !hasNext {
		return -1
	}
	car := entity.Car{
		CarNumber: carNumber,
		Color:     color,
	}
	p.slotService.BlockSlot(car, nextSlotId)
	return nextSlotId
}

func (p *ParkingService) ReleaseParkingSpace(carNumber string) {
	slotId, _ := p.slotService.GetCarSlot(carNumber)
	p.ReleaseParkingSpaceBySlotId(slotId)
}

func NewParkingService(slotService interfaces.ISlotService) interfaces.IParkingService {
	return &ParkingService{
		slotService,
	}
}
