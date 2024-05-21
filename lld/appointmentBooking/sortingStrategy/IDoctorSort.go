package sortingStrategy

import (
	"awesomeProject/lld/appointmentBooking/models"
)

type ISort interface {
	Sort([]models.Doctor) []models.Doctor
}
