package implement

import (
	"azizdev/helper"
	"azizdev/model/domain"
	"azizdev/repository"
	"context"
	"database/sql"
	"errors"
)

type EmployeeRepositoryImpl struct{}

func NewEmployeeRepository() repository.EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

func (repository *EmployeeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "insert into employee (first_name, last_name, id_card, hire_date, termination_date, salary) values (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, employee.FirstName, employee.LastName, employee.IdCard, employee.HireDate, employee.TerminationDate, employee.Salary)
	helper.PanicIfError(err)

	employeeId, err := result.LastInsertId()
	helper.PanicIfError(err)

	employee.EmployeeId = int(employeeId)
	return employee
}

func (repository *EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "update employee set first_name =?, last_name=?, id_card=?, hire_date=?, termination_date=?, salary=? where employee_id=?"
	_, err := tx.ExecContext(ctx, SQL, employee.FirstName, employee.LastName, employee.IdCard, employee.HireDate, employee.TerminationDate, employee.Salary)
	helper.PanicIfError(err)

	return employee
}

func (repository *EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee) {
	SQL := "delete from employee where employee_id=?"
	_, err := tx.ExecContext(ctx, SQL, employee.EmployeeId)
	helper.PanicIfError(err)
}

func (repository *EmployeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, employeeId int) (domain.Employee, error) {
	SQL := "select employee_id, first_name, last_name, id_card, hire_date, termination_date, salary from employee where employee_id=?"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
	defer rows.Close()

	employee := domain.Employee{}
	if rows.Next() {
		rows.Scan(&employee.EmployeeId, &employee.FirstName, &employee.LastName, &employee.IdCard, &employee.HireDate, &employee.TerminationDate, &employee.Salary)
		helper.PanicIfError(err)
		return employee, nil
	} else {
		return employee, errors.New("employee not found")
	}
}

func (repository *EmployeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee {
	SQL := `select employee_id, first_name, last_name, id_card, hire_date, termination_date, salary from employee`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var employees []domain.Employee
	for rows.Next() {
		employee := domain.Employee{}
		err := rows.Scan(&employee.EmployeeId, &employee.FirstName, &employee.LastName, &employee.IdCard, &employee.HireDate, &employee.TerminationDate, &employee.Salary)
		helper.PanicIfError(err)
		employees = append(employees, employee)
	}
	return employees
}
