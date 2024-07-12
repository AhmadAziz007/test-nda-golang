package domain

import "time"

type Employee struct {
	EmployeeId      int
	FirstName       string
	LastName        string
	IdCard          int
	HireDate        time.Time
	TerminationDate time.Time
	Salary          float64
}
