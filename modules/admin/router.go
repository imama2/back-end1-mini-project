package admin

import (
	"back-end1-mini-project/middleware"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
	adminHandler *AdminRequestHandler
}

//func NewAdminRouter(db *gorm.DB) *AdminRouter {
//	adminRepo := repositories.NewAdminRepository(db)
//	adminUsecase := AdminUseCase{
//		adminRepo: adminRepo,
//	}
//	adminController := AdminController{
//		uc: adminUsecase,
//	}
//	adminHandler := NewAdminRequestHandler(&adminController)
//
//	return &AdminRouter{
//		adminHandler: adminHandler,
//	}
//}

func (r *AdminRouter) Handle(router *gin.Engine) {
	basePath := "/account-admin"

	admin := router.Group(basePath)
	{
		admin.POST("/login", r.adminHandler.LoginAdmin)
		admin.POST("/register", r.adminHandler.RegisterAdmin)

		admin.Use(middleware.Authentication())
		{
			admin.POST("/create-customer", r.adminHandler.CreateCustomer)
			admin.DELETE("/delete-customer/:id", r.adminHandler.DeleteCustomerById)
			admin.GET("/customers", r.adminHandler.GetAllCustomers)
			admin.GET("/fetch-customers", r.adminHandler.SaveCustomersFromAPI)
		}
	}
}
