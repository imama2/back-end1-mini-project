package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminRequestHandler struct {
	adminController *AdminController
}

func NewAdminRequestHandler(adminController *AdminController) *AdminRequestHandler {
	return &AdminRequestHandler{
		adminController: adminController,
	}
}

func (handler *AdminRequestHandler) LoginAdmin(c *gin.Context) {
	request := AdminParam{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.adminController.LoginAdmin(request.ID, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AdminRequestHandler) RegisterAdmin(c *gin.Context) {
	request := AdminParam{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.adminController.RegisterAdmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AdminRequestHandler) CreateCustomer(c *gin.Context) {
	request := UserParam{}

	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := handler.adminController.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AdminRequestHandler) DeleteCustomerById(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	customerID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.adminController.DeleteUserById(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete Customer Data Successfully",
	})
}

func (handler *AdminRequestHandler) GetAllCustomers(c *gin.Context) {
	firstName := c.Query("first_name")
	lastName := c.Query("last_name")
	email := c.Query("email")

	// Parse pagination
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	response, err := handler.adminController.GetAllUsers(firstName, lastName, email, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (handler *AdminRequestHandler) SaveCustomersFromAPI(c *gin.Context) {
	response, err := handler.adminController.SaveUsersFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
