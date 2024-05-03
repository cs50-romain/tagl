package types

import "time"

type EmployeeItems struct {
	Id		int
	EmployeeName	string
	ItemName	string
	AcquisitionDate	time.Time
	Quantity	int
	TicketNumber	int
}

func NewEmployeeItems(employeeName, itemName string, quantity, ticketNum int, date time.Time) *EmployeeItems {
	return &EmployeeItems{
		EmployeeName: employeeName,
		ItemName: itemName,
		AcquisitionDate: date,
		Quantity: quantity,
		TicketNumber: ticketNum,
	}
}
