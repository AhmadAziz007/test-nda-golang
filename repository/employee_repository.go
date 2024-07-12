package repository

import (
	"azizdev/model/domain"
	"context"
	"database/sql"
)

type EmployeeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee)
	FindById(ctx context.Context, tx *sql.Tx, employeeId int) (domain.Employee, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Employee
}
