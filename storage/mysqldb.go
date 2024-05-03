package storage

import (
	"cs50-romain/tagl/types"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	Db *sql.DB	
}

func NewMysqlStore() (*MySQLDB, error) {
	user := "tagldb"
	password := "taglpass"
	dbname := "mysql"
	//ssldisable := true

	connStr := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	mydb := &MySQLDB{Db: db}
	err = mydb.createTable("inventory")
	if err != nil {
		return nil, err
	}

	return mydb, nil
}

func (s *MySQLDB) GetEmployeeByID(id int) *types.EmployeeItems {
	return &types.EmployeeItems{
		Id: id,
		EmployeeName: "name",
	}	
}

func (s *MySQLDB) CreateEmployee(employee *types.EmployeeItems) error {
	result, err := s.Db.Exec("insert into EmployeeItems (employee_name, item_name, acquisition_date, quantity, ticket_number) values (?, ?, ?, ?, ?)", employee.EmployeeName, employee.ItemName, time.Now(), employee.Quantity, employee.TicketNumber)

	if err != nil {
		fmt.Println("error:", err)
		return err 
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (s *MySQLDB) createTable(tablename string) error {
	query := `create table if not exists EmployeeItems (
		    record_id SERIAL PRIMARY KEY,
		    employee_name VARCHAR(100) NOT NULL,
		    item_name VARCHAR(100) NOT NULL,
		    acquisition_date DATE NOT NULL,
		    quantity INT NOT NULL,
	            ticket_number VARCHAR(100)
		);`

	_, err := s.Db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
