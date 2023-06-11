package customer

import (
	"back-end1-mini-project/dto"
	"back-end1-mini-project/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CustomerRequestHandler struct {
	controllerCustomer CustomerController
}

func NewCustomerRequestHandler(db *gorm.DB) CustomerRequestHandler {
	customerRepo := repositories.NewCustomerRepository(db)
	uc := CustomerUseCase{customerRepo}
	controllerCustomer := CustomerController{uc}

	return CustomerRequestHandler{controllerCustomer}
}

func (rh CustomerRequestHandler) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := rh.controllerCustomer.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, res)
}

func (rh CustomerRequestHandler) GetCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.controllerCustomer.GetCustomerById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh CustomerRequestHandler) UpdateCustomerById(c *gin.Context) {
	id := c.Param("id")
	request := CustomerParam{}

	// Bind JSON
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.controllerCustomer.UpdateCustomerById(uint(customerID), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh CustomerRequestHandler) DeleteCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	err = rh.controllerCustomer.DeleteCustomerById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete Customer Data Successfully",
	})
}
