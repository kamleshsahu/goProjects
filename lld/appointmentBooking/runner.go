package appointmentBooking

import (
	"awesomeProject/lld/appointmentBooking/appointmentManager"
	"awesomeProject/lld/appointmentBooking/doctor"
	models2 "awesomeProject/lld/appointmentBooking/models"
	"awesomeProject/lld/appointmentBooking/patient"
	"awesomeProject/lld/appointmentBooking/sortingStrategy"
	"fmt"
)

func Runner() {
	ds := doctor.New(sortingStrategy.Default())
	ps := patient.New()

	am := appointmentManager.New(ds, ps)
	user := models2.User{
		Name: "Kamlesh",
	}
	_doctor := models2.Doctor{
		User:       user,
		Speciality: Cardiologist,
		SlotMap:    make(map[int]int),
		Rating:     7,
	}

	doctor1 := am.RegisterDoctor(_doctor)

	_ = am.RegisterDoctor(models2.Doctor{
		User: models2.User{
			Name: "Nilesh",
		},
		Speciality: Dermatologist,
		SlotMap:    make(map[int]int),
		Rating:     3,
	})

	_ = am.RegisterDoctor(models2.Doctor{
		User: models2.User{
			Name: "Tikesh",
		},
		Speciality: Dermatologist,
		SlotMap:    make(map[int]int),
		Rating:     2,
	})

	_patient := models2.Patient{
		User:     user,
		Bookings: make(map[int]int),
	}
	patient1 := am.RegisterUser(_patient)
	am.AddSlot(doctor1, 1)
	am.AddSlot(doctor1, 2)
	am.AddSlot(doctor1, 4)
	am.BookAppointment(doctor1, patient1, 1)
	am.BookAppointment(doctor1, patient1, 2)

	patient2 := am.RegisterUser(models2.Patient{
		User:     user,
		Bookings: make(map[int]int),
	})

	am.BookAppointment(doctor1, patient2, 3)
	am.CancelAppointment(doctor1, patient1, 1)
	am.DeleteSlot(doctor1, 1)
	fmt.Println(am.GetAppointments(doctor1))
	fmt.Println(am.GetMyAppointments(patient1))
	fmt.Println(am.GetMyAppointments(patient2))

	list := ds.GetDoctors()

	fmt.Println(list)
	for _, data := range list {
		fmt.Println(data.User.Name, data.Rating)
	}
}
