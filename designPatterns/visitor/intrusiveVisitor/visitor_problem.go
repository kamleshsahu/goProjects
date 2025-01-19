package main

import "fmt"

type Expression interface {
	Print()
}

type DoubleExpression struct {
	value float64
}

func (d DoubleExpression) Print() {
	fmt.Print(d.value)
}

type AdditionExpression struct {
	left, right Expression
}

func (a AdditionExpression) Print() {
	a.left.Print()
	a.right.Print()
}

// (1+2) + 3

func main() {
	// 1 + (2 + 3 )

	val := AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	val.Print()
}
