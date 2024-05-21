package service

import (
	"awesomeProject/lld/elevator/entity"
	"awesomeProject/lld/elevator/interfaces"
	"fmt"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type elevatorController struct {
	upRequests   *priorityqueue.Queue
	downRequests *priorityqueue.Queue
	elevatorCar  entity.ElevatorCar
}

func (e *elevatorController) Move() {

	if e.elevatorCar.Direction == entity.DOWN {
		if e.downRequests.Empty() {
			e.elevatorCar.ElevatorState = entity.STILL
			e.ChangeDirection(entity.UP)
			return
		}
		nextFloor, _ := e.downRequests.Dequeue()
		e.elevatorCar.ElevatorState = entity.MOVING

		e.elevatorCar.CurrentFloor = nextFloor.(int)
	} else {
		if e.upRequests.Empty() {
			e.elevatorCar.ElevatorState = entity.STILL
			e.ChangeDirection(entity.DOWN)
			return
		}
		nextFloor, _ := e.upRequests.Dequeue()
		e.elevatorCar.ElevatorState = entity.MOVING

		e.elevatorCar.CurrentFloor = nextFloor.(int)
	}
	fmt.Println("reached floor", e.elevatorCar.CurrentFloor)
	fmt.Println("moving to next")
	e.Move()
}

func (e *elevatorController) ChangeDirection(direction entity.Direction) {
	if e.elevatorCar.Direction == direction {
		return
	}

	if e.upRequests.Empty() && e.downRequests.Empty() {
		return
	}

	e.elevatorCar.Direction = direction
	e.Start()
}

func (e *elevatorController) Start() {
	if e.elevatorCar.ElevatorState == entity.STILL {
		e.Move()
	}
}

func (e *elevatorController) SubmitExternalRequest(floor int, direction entity.Direction) {

	if e.elevatorCar.CurrentFloor > floor && direction == entity.DOWN {
		e.downRequests.Enqueue(floor)
	} else {
		e.upRequests.Enqueue(floor)
	}
	e.Start()
}

func (e *elevatorController) SubmitInternalRequest(toFloor int) {

	if e.elevatorCar.CurrentFloor > toFloor {
		e.downRequests.Enqueue(toFloor)
	} else {
		e.upRequests.Enqueue(toFloor)
	}

	e.Start()
}

func NewEC(car entity.ElevatorCar) interfaces.IElevatorController {
	return &elevatorController{elevatorCar: car,
		upRequests: priorityqueue.NewWith(utils.IntComparator),
		downRequests: priorityqueue.NewWith(func(a, b interface{}) int {
			return -utils.IntComparator(a, b)
		}),
	}
}
