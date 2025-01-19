package factories

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// functional
func NewEmployeeFactory(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 60000}
	case Manager:
		return &Employee{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}

func main() {
	m := NewEmployeeFactory(Manager)
	m.Name = "Sam"

	dev := NewEmployeeFactory(Developer)
	dev.Name = "SamDeveloper"
	fmt.Println(dev)
}
