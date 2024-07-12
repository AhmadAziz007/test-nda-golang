package main

import (
	"azizdev/app"
	"azizdev/controller/implement"
	"azizdev/helper"
	"azizdev/middleware"
	implement2 "azizdev/repository/implement"
	implement3 "azizdev/service/implement"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	employeeRepository := implement2.NewEmployeeRepository()
	employeeService := implement3.NewEmployeeService(employeeRepository, validate, db)
	employeeController := implement.NewEmployeeController(employeeService)

	router := app.NewRouter(employeeController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
