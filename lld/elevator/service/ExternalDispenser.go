package service

import (
	"awesomeProject/lld/elevator/entity"
	"awesomeProject/lld/elevator/interfaces"
)

type externalDispenser struct {
	CurrentFloor int
	ecList       []interfaces.IElevatorController
}

func (e *externalDispenser) PressButton(direction entity.Direction) {
	e.SubmitExternalRequest(direction)
}

func (e *externalDispenser) SubmitExternalRequest(direction entity.Direction) {
	e.ecList[0].SubmitExternalRequest(e.CurrentFloor, direction)
}

func NewExternalDispenser(currFloor int, ecList []interfaces.IElevatorController) interfaces.IExternalDispenser {
	return &externalDispenser{CurrentFloor: currFloor, ecList: ecList}
}
