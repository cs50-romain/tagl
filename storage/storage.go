package storage

import "cs50-romain/tagl/types"

type Storer interface {
	createTable(string) error
	GetEmployeeByID(int) *types.EmployeeItems
	CreateEmployee(*types.EmployeeItems) error
}
