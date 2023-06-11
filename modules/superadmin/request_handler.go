package superadmin

import (
	"back-end1-mini-project/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RequestHandlerSuperadmin struct {
	ctrl ControllerSuperadminInterface
}

func NewRequestHandlerSuperadmin(ctrl ControllerSuperadminInterface) *RequestHandlerSuperadmin {
	return &RequestHandlerSuperadmin{
		ctrl: ctrl,
	}
}

func (rh *RequestHandlerSuperadmin) CreateSuperadmin(c *gin.Context) {
	request := SuperAdminParam{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.CreateSuperadmin(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) LoginSuperadmin(c *gin.Context) {
	request := SuperAdminParam{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.LoginSuperadmin(request.ID, request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) CreateUser(c *gin.Context) {
	request := CustomerParam{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	userID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	err = rh.ctrl.DeleteUserById(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete Customer Data Successfully",
	})
}

func (rh *RequestHandlerSuperadmin) GetAllUsers(c *gin.Context) {
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

	response, err := rh.ctrl.GetAllUsers(firstName, lastName, email, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) ApproveAdminRegistration(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	approvedID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.ApproveAdminRegistration(uint(approvedID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) RejectAdminRegistration(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	rejectedID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.RejectAdminRegistration(uint(rejectedID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) UpdateAdminActiveStatus(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	adminID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.UpdateAdminActiveStatus(uint(adminID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) UpdateAdminInactiveStatus(c *gin.Context) {
	id := c.Param("id")

	// Parse id to uint
	adminID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	response, err := rh.ctrl.UpdateAdminInactiveStatus(uint(adminID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) GetApprovalRequests(c *gin.Context) {
	response, err := rh.ctrl.GetApprovalRequests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}

func (rh *RequestHandlerSuperadmin) GetAllAdmins(c *gin.Context) {
	username := c.Query("username")

	// Parse pagination
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		pageSize = 10
	}

	response, err := rh.ctrl.GetAllAdmins(username, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}

	c.JSON(http.StatusOK, response)
}
