package app

import (
	"azizdev/controller"
	"azizdev/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(employeeController controller.EmployeeController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/employees", employeeController.FindAll)
	router.GET("/api/employee/:employeeId", employeeController.FindById)
	router.POST("/api/employee", employeeController.Create)
	router.PUT("/api/employee/:employeeId", employeeController.Update)
	router.DELETE("/api/employee/:employeeId", employeeController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
