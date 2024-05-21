package service

import "awesomeProject/calculator/interfaces"

type number struct {
	Value int
}

func (e *number) Evaluate() int {
	return e.Value
}

func Number(value int) interfaces.IExpression {
	return &number{
		Value: value,
	}
}
