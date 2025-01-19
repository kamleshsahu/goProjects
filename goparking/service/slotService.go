package service

import (
	"awesomeProject/goparking/entity"
	"awesomeProject/goparking/interfaces"
)

type SlotService struct {
	size  int
	slots []entity.Slot
}

func (s *SlotService) GetAllCarsByColor(color string) []entity.Car {
	cars := make([]entity.Car, 0)
	for _, slot := range s.slots {
		if slot.IsBlocked {
			if slot.Car.Color == color {
				cars = append(cars, *slot.Car)
			}
		}
	}
	return cars
}

func (s *SlotService) GetCarSlot(carNumber string) (int, bool) {
	for i, slot := range s.slots {
		if slot.IsBlocked {
			if slot.Car.CarNumber == carNumber {
				return i, true
			}
		}
	}
	return -1, false
}

func (s *SlotService) GetNextSlot() (int, bool) {
	for i, slot := range s.slots {
		if !slot.IsBlocked {
			return i, true
		}
	}
	return -1, false
}

func (s *SlotService) BlockSlot(car entity.Car, slotId int) bool {
	slot := s.slots[slotId]
	if slot.IsBlocked {
		return false
	}
	slot.IsBlocked = true
	slot.Car = &car
	s.slots[slotId] = slot
	return true
}

func (s *SlotService) ReleaseSlot(slotId int) bool {
	slot := s.slots[slotId]
	if !slot.IsBlocked {
		return true
	}
	slot.IsBlocked = false
	slot.Car = nil
	s.slots[slotId] = slot
	return true
}

func NewSlotService(n int) interfaces.ISlotService {
	slots := make([]entity.Slot, n)

	return &SlotService{
		size:  n,
		slots: slots,
	}
}
