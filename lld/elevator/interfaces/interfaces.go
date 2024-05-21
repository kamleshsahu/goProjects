package interfaces

import "awesomeProject/lld/elevator/entity"

type IExternalDispenser interface {
	PressButton(direction entity.Direction)
	SubmitExternalRequest(direction entity.Direction)
}

type IInternalDispenser interface {
	SubmitInternalRequest(toFloor int)
}

type IElevatorController interface {
	SubmitExternalRequest(toFloor int, direction entity.Direction)
	SubmitInternalRequest(toFloor int)
	Move()
}
