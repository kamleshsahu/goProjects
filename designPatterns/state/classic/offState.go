package main

import "fmt"

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on...")
	sw.State = NewOnState()
}

type OffState struct {
	BaseState
}
