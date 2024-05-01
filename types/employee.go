package types

type Employee struct {
	Id	int
	Name	string
}

func NewEmployee(name string) *Employee {
	return &Employee{
		Name: name,
	}
}
