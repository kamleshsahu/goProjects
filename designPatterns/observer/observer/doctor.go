package observer

import "fmt"

type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}
