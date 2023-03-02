package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponseM struct {
	Message string `json:"message"`
}

type HandlerM struct {
	managment Managment
}

func NewHandlerM(managment Managment) *HandlerM {
	return &HandlerM{managment: managment}
}

func (hm *HandlerM) CreateDepartment(c *gin.Context) {
	var department Department
	
	if err := c.BindJSON(&department); err != nil {
		fmt.Printf("failed to bind department: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponseM{
			Message: err.Error(),
		})
		return
	}

	hm.managment.InsertM(&department)

	c.JSON(http.StatusOK, map[string]interface{}{
		"office": department.Office,
		
	})
}

func (hm *HandlerM) GetDepartment(c *gin.Context) {
	office, err := strconv.Atoi(c.Param("office"))
	if err != nil {
		fmt.Printf("failed to convert name param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponseM{
			Message: err.Error(),
		})
		return
	}

	department, err := hm.managment.GetM(office)
	if err != nil {
		fmt.Printf("failed to get department %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponseM{
			Message: err.Error(),
		})
		return
	}
	
	
	c.JSON(http.StatusOK, department)
}

func (hm *HandlerM) UpdateDepartment(c *gin.Context) {
	office, err := strconv.Atoi(c.Param("office"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var department Department

	if err := c.BindJSON(&department); err != nil {
		fmt.Printf("failed to bind department: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	hm.managment.UpdateM(office, department)

	c.JSON(http.StatusOK, map[string]interface{}{
		"office": department.Office,
	})
}

func (hm *HandlerM) DeleteDepartment(c *gin.Context) {
	office, err := strconv.Atoi(c.Param("office"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	hm.managment.DeleteM(office)

	c.String(http.StatusOK, "department deleted")
}