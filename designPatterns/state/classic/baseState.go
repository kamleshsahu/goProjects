package main

import "fmt"

type BaseState struct{}

func (s *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}
