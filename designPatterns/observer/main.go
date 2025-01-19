package main

import (
	"awesomeProject/designPatterns/observer/observable"
	"awesomeProject/designPatterns/observer/observer"
)

func main() {
	p := observable.NewPerson("Boris")
	ds := &observer.DoctorService{}
	p.Subscribe(ds)

	is := &observer.InsService{}
	p.Subscribe(is)

	// let's test it!
	p.CatchACold()
}
