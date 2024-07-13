package response

import "time"

type EmployeeResponse struct {
	EmployeeId      int       `json:"employeeId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	IdCard          int       `json:"idCard"`
	HireDate        time.Time `json:"hireDate"`
	TerminationDate time.Time `json:"terminationDate"`
	Salary          float64   `json:"salary"`
}
