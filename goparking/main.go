package goparking

import (
	"awesomeProject/goparking/service"
	"fmt"
)

func Run() {

	n := 5
	ss := service.NewSlotService(n)
	ps := service.NewParkingService(ss)

	blueCars := ps.GetAllCarsByColor("blue")

	fmt.Println(blueCars)
	ps.AllotParkingSpace("123", "blue")
	blueCars = ps.GetAllCarsByColor("blue")
	fmt.Println(blueCars)
	ps.ReleaseParkingSpace("123")
	blueCars = ps.GetAllCarsByColor("blue")
	fmt.Println(blueCars)
	ps.AllotParkingSpace("124", "red")
	ps.AllotParkingSpace("125", "blue")
	ps.AllotParkingSpace("126", "blue")
	ps.AllotParkingSpace("127", "blue")
	ps.ReleaseParkingSpace("126")
	blueCars = ps.GetAllCarsByColor("blue")
	fmt.Println(blueCars)
}
