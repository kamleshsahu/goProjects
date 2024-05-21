package calculator

import (
	"awesomeProject/calculator/interfaces"
	"awesomeProject/calculator/service"
	"fmt"
)

func Run() {

	se := service.ArithmeticExp(service.ArithmeticExp(service.Number(150), service.Number(10), interfaces.DIVIDE), service.ArithmeticExp(service.Number(15), service.Number(10), interfaces.DIVIDE), interfaces.DIVIDE) // = 1

	fmt.Println(se.Evaluate())
}
