package model

import (
	"azizdev/model/domain"
	"azizdev/model/web/response"
)

func ToEmployeeResponse(employee domain.Employee) response.EmployeeResponse {
	return response.EmployeeResponse{
		EmployeeId:      employee.EmployeeId,
		FirstName:       employee.FirstName,
		LastName:        employee.LastName,
		IdCard:          employee.IdCard,
		HireDate:        employee.HireDate,
		TerminationDate: employee.TerminationDate,
		Salary:          employee.Salary,
	}
}

func ToEmployeeResponses(employees []domain.Employee) []response.EmployeeResponse {
	var employeeResponses []response.EmployeeResponse
	for _, employeeList := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employeeList))
	}
	return employeeResponses
}
