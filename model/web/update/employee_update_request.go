package update

import "time"

type EmployeeUpdateRequest struct {
	EmployeeId      int       `validate:"required"`
	FirstName       string    `validate:"required,min=1,max=100" json:"firstName"`
	LastName        string    `validate:"required,min=1,max=100" json:"lastName"`
	IdCard          int       `validate:"required" json:"idCard"`
	HireDate        time.Time `validate:"required" json:"hireDate"`
	TerminationDate time.Time `validate:"required" json:"terminationDate"`
	Salary          float64   `validate:"required" json:"salary"`
}
