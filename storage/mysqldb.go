package storage

import (
	"cs50-romain/tagl/types"
	"database/sql"
	"errors"
)

type MySQLDB struct {
	db *sql.DB	
}

func NewMysqlStore() (*MySQLDB, error) {
	return nil, errors.New("cannot open yet")
}

func (s *MySQLDB) Get(id int) *types.Employee {
	return &types.Employee{
		Id: id,
		Name: "name",
	}	
}
