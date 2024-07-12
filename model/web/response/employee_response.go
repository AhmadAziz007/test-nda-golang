package response

import "time"

type EmployeeResponse struct {
	EmployeeId      int       `json:"employee_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	IdCard          int       `json:"id_card"`
	HireDate        time.Time `json:"hire_date"`
	TerminationDate time.Time `json:"termination_date"`
	Salary          float64   `json:"salary"`
}
