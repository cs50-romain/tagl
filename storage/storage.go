package storage

import "cs50-romain/tagl/types"

type Storer interface {
	createTable(string) error
	GetEmployeeByName(string) ([]*types.EmployeeItems, error)
	CreateEmployee(*types.EmployeeItems) error
}
