package service

import "awesomeProject/lld/elevator/interfaces"

type internalDispenser struct {
	elevatorController interfaces.IElevatorController
}

func (i *internalDispenser) SubmitInternalRequest(toFloor int) {
	i.elevatorController.SubmitInternalRequest(toFloor)
}

func NewInternalDispenser(controller interfaces.IElevatorController) interfaces.IInternalDispenser {
	return &internalDispenser{elevatorController: controller}
}
