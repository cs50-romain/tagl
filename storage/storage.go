package storage

import "cs50-romain/tagl/types"

type Storer interface {
	Get(int) *types.Employee
}
