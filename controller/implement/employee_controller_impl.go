package implement

import (
	"azizdev/controller"
	"azizdev/helper"
	"azizdev/model/web"
	"azizdev/model/web/create"
	"azizdev/model/web/update"
	"azizdev/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type EmployeeControllerImpl struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(employeeService service.EmployeeService) controller.EmployeeController {
	return &EmployeeControllerImpl{
		EmployeeService: employeeService}
}

func (controller *EmployeeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeCreateRequest := create.EmployeeCreateRequest{}
	helper.ReadFromRequestBody(request, &employeeCreateRequest)

	employeeResponse := controller.EmployeeService.Create(request.Context(), employeeCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EmployeeControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeUpdateRequest := update.EmployeeUpdateRequest{}
	helper.ReadFromRequestBody(request, &employeeUpdateRequest)

	employeeId := params.ByName("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.PanicIfError(err)

	employeeUpdateRequest.EmployeeId = id

	employeeResponse := controller.EmployeeService.Update(request.Context(), employeeUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EmployeeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeId := params.ByName("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.PanicIfError(err)

	controller.EmployeeService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EmployeeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeId := params.ByName("employeeId")
	id, err := strconv.Atoi(employeeId)
	helper.PanicIfError(err)

	employeeResponse := controller.EmployeeService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *EmployeeControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	employeeResponses := controller.EmployeeService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   employeeResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
