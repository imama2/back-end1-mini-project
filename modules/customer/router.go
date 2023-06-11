package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomerRouter struct {
	RequestHandlerCustomer CustomerRequestHandler
}

func NewCustomerRouter(db *gorm.DB) CustomerRouter {
	return CustomerRouter{
		RequestHandlerCustomer: NewCustomerRequestHandler(db),
	}
}

func (r CustomerRouter) Handle(engine *gin.Engine) {
	basePath := "/Customer"

	customer := engine.Group(basePath)

	customer.POST("/create", r.RequestHandlerCustomer.CreateCustomer)
	customer.GET("/:id", r.RequestHandlerCustomer.GetCustomerById)
	customer.PUT("/:id", r.RequestHandlerCustomer.UpdateCustomerById)
	customer.DELETE("/:id", r.RequestHandlerCustomer.DeleteCustomerById)
}
