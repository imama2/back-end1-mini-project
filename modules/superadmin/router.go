package superadmin

import (
	"back-end1-mini-project/middleware"
	"github.com/gin-gonic/gin"
)

type RouterSuperadmin struct {
	SuperadminRequestHandler *RequestHandlerSuperadmin
}

func NewRouterSuperadmin(ctrl ControllerSuperadminInterface) *RouterSuperadmin {
	return &RouterSuperadmin{
		SuperadminRequestHandler: NewRequestHandlerSuperadmin(ctrl),
	}
}

func (r *RouterSuperadmin) Handle(router *gin.Engine) {
	basePath := "/account-superadmin"

	superadmin := router.Group(basePath)
	superadmin.POST("/register", r.SuperadminRequestHandler.CreateSuperadmin)
	superadmin.POST("/login", r.SuperadminRequestHandler.LoginSuperadmin)

	// About User
	superadmin.Use(middleware.Authentication())
	superadmin.GET("/customers", r.SuperadminRequestHandler.GetAllUsers)
	superadmin.POST("/create-customer", r.SuperadminRequestHandler.CreateUser)
	superadmin.DELETE("/delete-customer/:id", r.SuperadminRequestHandler.DeleteUserByID)

	// About Admin
	superadmin.Use(middleware.Authentication())
	superadmin.GET("/accounts", r.SuperadminRequestHandler.GetAllAdmins)
	superadmin.POST("/:id/approved", r.SuperadminRequestHandler.ApproveAdminRegistration)
	superadmin.POST("/:id/rejected", r.SuperadminRequestHandler.RejectAdminRegistration)
	superadmin.POST("/:id/actived", r.SuperadminRequestHandler.UpdateAdminActiveStatus)
	superadmin.POST("/:id/deadactived", r.SuperadminRequestHandler.UpdateAdminInactiveStatus)
	superadmin.GET("/approval-request", r.SuperadminRequestHandler.GetApprovalRequests)
}
