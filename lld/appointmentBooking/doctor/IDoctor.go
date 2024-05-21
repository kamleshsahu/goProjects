package doctor

import (
	"awesomeProject/lld/appointmentBooking/models"
	"awesomeProject/lld/appointmentBooking/sortingStrategy"
	"awesomeProject/lld/appointmentBooking/user"
)

type IDoctor interface {
	user.IUser
	AddSlot(doctorId, slotId int) bool
	DeleteSlot(doctorId, slotId int) bool
	UpdateSlot(doctorId, slotId, userId int) bool
	GetDoctors() []models.Doctor
	GetAppointments(doctorId int) map[int]int
}

type DoctorService struct {
	doctorMap       map[int]models.Doctor
	id              int
	SortingStrategy sortingStrategy.ISort
}

func (ds *DoctorService) GetDoctors() []models.Doctor {
	var dl []models.Doctor
	for _, value := range ds.doctorMap {
		dl = append(dl, value)
	}
	return ds.SortingStrategy.Sort(dl)
}

func (ds *DoctorService) GetAppointments(doctorId int) map[int]int {
	return ds.doctorMap[doctorId].SlotMap
}

func (ds *DoctorService) UpdateSlot(doctorId, slotId int, userId int) bool {
	ds.doctorMap[doctorId].SlotMap[slotId] = userId
	return true
}

func (ds *DoctorService) AddSlot(doctorId, slotId int) bool {
	ds.doctorMap[doctorId].SlotMap[slotId] = 0
	return true
}

func (ds *DoctorService) DeleteSlot(doctorId, slotId int) bool {
	delete(ds.doctorMap[doctorId].SlotMap, slotId)
	return true
}

func (ds *DoctorService) Register(doctor interface{}) int {
	doctorId := ds.id
	ds.id++
	ds.doctorMap[doctorId] = doctor.(models.Doctor)

	return doctorId
}

func New(sort sortingStrategy.ISort) IDoctor {
	return &DoctorService{
		doctorMap:       make(map[int]models.Doctor),
		id:              1,
		SortingStrategy: sort,
	}
}
