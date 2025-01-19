package observer

import "fmt"

type InsService struct{}

func (d *InsService) Notify(data interface{}) {
	fmt.Printf("\nA insurance company has been called for %s",
		data.(string))
}
