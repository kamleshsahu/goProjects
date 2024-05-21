package patient

import (
	"awesomeProject/lld/appointmentBooking/models"
	"awesomeProject/lld/appointmentBooking/user"
)

type PatientService struct {
	patientMap map[int]models.Patient
	id         int
}

func (p *PatientService) Register(patient any) int {
	patientId := p.id
	p.id++
	p.patientMap[patientId] = patient.(models.Patient)

	return patientId
}

func (p *PatientService) BookAppointment(doctorId, userId, slotId int) bool {
	p.patientMap[userId].Bookings[slotId] = doctorId
	return true
}

func (p *PatientService) CancelAppointment(userId, slotId int) bool {
	delete(p.patientMap[userId].Bookings, slotId)
	return true
}

func (p *PatientService) GetMyAppointments(userId int) map[int]int {
	return p.patientMap[userId].Bookings
}

func New() IPatient {
	return &PatientService{
		patientMap: make(map[int]models.Patient),
		id:         1,
	}
}

type IPatient interface {
	user.IUser
	BookAppointment(doctorId, userId, slotId int) bool
	CancelAppointment(userId, slotId int) bool
	GetMyAppointments(userId int) map[int]int
}
