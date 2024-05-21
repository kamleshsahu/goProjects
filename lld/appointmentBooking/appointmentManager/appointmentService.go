package appointmentManager

import (
	"awesomeProject/lld/appointmentBooking/doctor"
	"awesomeProject/lld/appointmentBooking/patient"
)

type IAppointmentService interface {
	RegisterUser(user interface{}) int
	RegisterDoctor(doctor interface{}) int
	BookAppointment(doctorId, userId, slotId int) bool
	CancelAppointment(doctorId, userId, slotId int) bool
	AddSlot(doctorId, slotId int) bool
	DeleteSlot(doctorId, slotId int) bool
	GetAppointments(doctorId int) map[int]int
	GetMyAppointments(userId int) map[int]int
}

type AppointmentService struct {
	doctor  doctor.IDoctor
	patient patient.IPatient
}

func (as *AppointmentService) GetMyAppointments(userId int) map[int]int {
	return as.patient.GetMyAppointments(userId)
}

func (as *AppointmentService) GetAppointments(doctorId int) map[int]int {
	return as.doctor.GetAppointments(doctorId)
}

func (as *AppointmentService) RegisterUser(user interface{}) int {
	return as.patient.Register(user)
}

func (as *AppointmentService) RegisterDoctor(doctor interface{}) int {
	return as.doctor.Register(doctor)
}

func (as *AppointmentService) CancelAppointment(doctorId, userId, slotId int) bool {
	as.doctor.UpdateSlot(doctorId, slotId, 0)
	as.patient.CancelAppointment(userId, slotId)
	return true
}

func (as *AppointmentService) AddSlot(doctorId, slotId int) bool {
	return as.doctor.AddSlot(doctorId, slotId)
}

func (as *AppointmentService) DeleteSlot(doctorId, slotId int) bool {
	as.doctor.DeleteSlot(doctorId, slotId)
	return true
}

func (as *AppointmentService) BookAppointment(doctorId, userId, slotId int) bool {

	as.doctor.UpdateSlot(doctorId, slotId, userId)
	as.patient.BookAppointment(doctorId, userId, slotId)
	return true
}

func New(doctor doctor.IDoctor, patient patient.IPatient) IAppointmentService {
	return &AppointmentService{
		doctor:  doctor,
		patient: patient,
	}
}
