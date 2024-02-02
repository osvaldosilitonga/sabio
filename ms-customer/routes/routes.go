package routes

import (
	"ms-customer/configs"
	"ms-customer/controllers"
	"ms-customer/repositories"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	customerRepository := repositories.NewCustomerRepository(db)

	customerController := controllers.NewCustomerController(customerRepository)

	// Routes
	v1 := e.Group("/v1")
	{
		v1.POST("/customer", customerController.Create)
		v1.POST("/customer", customerController.Update)
		v1.DELETE("/customer", customerController.Delete)
		v1.GET("/customer", customerController.FindAll)
		v1.GET("/customer/:id", customerController.FindById)
	}
}
