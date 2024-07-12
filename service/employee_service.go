package service

import (
	"azizdev/model/web/create"
	"azizdev/model/web/response"
	"azizdev/model/web/update"
	"context"
)

type EmployeeService interface {
	Create(ctx context.Context, request create.EmployeeCreateRequest) response.EmployeeResponse
	Update(ctx context.Context, request update.EmployeeUpdateRequest) response.EmployeeResponse
	Delete(ctx context.Context, EmployeeId int)
	FindById(ctx context.Context, EmployeeId int) response.EmployeeResponse
	FindAll(ctx context.Context) []response.EmployeeResponse
}
