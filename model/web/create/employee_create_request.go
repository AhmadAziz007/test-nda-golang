package create

import "time"

type EmployeeCreateRequest struct {
	FirstName       string    `validate:"required,min=1,max=100" json:"firstName"`
	LastName        string    `validate:"required,min=1,max=100" json:"lastName"`
	IdCard          int       `validate:"required" json:"idCard"`
	HireDate        time.Time `validate:"required"`
	TerminationDate time.Time `validate:"required"`
	Salary          float64   `validate:"required" json:"salary"`
}
