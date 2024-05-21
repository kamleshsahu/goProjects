package service

import "awesomeProject/calculator/interfaces"

type expression struct {
	Exp1      interfaces.IExpression
	Exp2      interfaces.IExpression
	Operation interfaces.Operation
}

func (e *expression) Evaluate() int {
	res1 := e.Exp1.Evaluate()
	res2 := e.Exp2.Evaluate()

	switch e.Operation {
	case interfaces.ADD:
		return res1 + res2
	case interfaces.SUBTRACT:
		return res1 - res2
	case interfaces.MULTIPLY:
		return res1 * res2
	case interfaces.DIVIDE:
		return res1 / res2
	}
	return 0
}

func ArithmeticExp(exp1, exp2 interfaces.IExpression, operation interfaces.Operation) interfaces.IExpression {
	return &expression{
		Exp1:      exp1,
		Exp2:      exp2,
		Operation: operation,
	}
}
