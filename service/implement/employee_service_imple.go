package implement

import (
	"azizdev/exception"
	"azizdev/helper"
	"azizdev/helper/model"
	"azizdev/model/domain"
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"azizdev/repository"
	"azizdev/service"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository, DB *validator.Validate, validate *sql.DB) service.EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		DB:                 validate,
		Validate:           DB}
}

func (service *EmployeeServiceImpl) Create(ctx context.Context, request create.EmployeeCreateRequest) response.EmployeeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee := domain.Employee{
		FirstName:       request.FirstName,
		LastName:        request.LastName,
		IdCard:          request.IdCard,
		HireDate:        request.HireDate,
		TerminationDate: request.TerminationDate,
		Salary:          request.Salary,
	}

	employee = service.EmployeeRepository.Save(ctx, tx, employee)

	return model.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) Update(ctx context.Context, request update.EmployeeUpdateRequest) response.EmployeeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, request.EmployeeId)
	helper.PanicIfError(err)

	employee.FirstName = request.FirstName
	employee.LastName = request.LastName
	employee.IdCard = request.IdCard
	employee.HireDate = request.HireDate
	employee.TerminationDate = request.TerminationDate
	employee.Salary = request.Salary

	employee = service.EmployeeRepository.Save(ctx, tx, employee)

	return model.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) Delete(ctx context.Context, EmployeeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, EmployeeId)
	helper.PanicIfError(err)
	service.EmployeeRepository.Delete(ctx, tx, employee)
}

func (service *EmployeeServiceImpl) FindById(ctx context.Context, EmployeeId int) response.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employee, err := service.EmployeeRepository.FindById(ctx, tx, EmployeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) FindAll(ctx context.Context) []response.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	employees := service.EmployeeRepository.FindAll(ctx, tx)
	return model.ToEmployeeResponses(employees)
}
