package main	

import (
	"github.com/gin-gonic/gin"
)


func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	memoryManagment := NewMemoryManagment()
	handlerM := NewHandlerM(memoryManagment)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	router.GET("/employee/all", handler.GetAllEmployee)
	router.GET("/employee/all/:office", handler.GetEmployeeOffice, handlerM.GetDepartment)
	
	
	router.POST("/department", handlerM.CreateDepartment)
	router.GET("/department/:office", handlerM.GetDepartment)
	router.PUT("/department/:office", handlerM.UpdateDepartment)
	router.DELETE("/department/:office", handlerM.DeleteDepartment)
	router.GET("/department/:office/:id", handlerM.GetDepartment, handler.GetEmployee)
	router.GET("/department/:office/all",  handlerM.GetDepartment, handler.GetEmployeeOffice)
	
	
	router.Run()
}