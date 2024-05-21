package elevator

import (
	"awesomeProject/lld/elevator/entity"
	"awesomeProject/lld/elevator/interfaces"
	"awesomeProject/lld/elevator/service"
)

func Runner() {

	car := entity.ElevatorCar{CurrentFloor: 0, Direction: entity.UP, ElevatorState: entity.STILL}
	ec := service.NewEC(car)

	eDispenser := service.NewExternalDispenser(5, []interfaces.IElevatorController{ec})

	eDispenser.PressButton(entity.UP)

	iDispenser := service.NewInternalDispenser(ec)

	iDispenser.SubmitInternalRequest(10)

	iDispenser.SubmitInternalRequest(12)

	iDispenser.SubmitInternalRequest(4)

	iDispenser.SubmitInternalRequest(6)
	iDispenser.SubmitInternalRequest(0)

}
