package interfaces

const (
	ADD = iota
	SUBTRACT
	MULTIPLY
	DIVIDE
)

type Operation int

type IExpression interface {
	Evaluate() int
}
