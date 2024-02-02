package routes

import (
	"ms-product/configs"
	"ms-product/controllers"
	"ms-product/repositories"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	productRepository := repositories.NewProductRepository(db)

	productController := controllers.NewProductController(productRepository)

	// Routes
	v1 := e.Group("/v1")
	{
		v1.POST("/product", productController.Create)
		v1.POST("/product", productController.Update)
		v1.DELETE("/product", productController.Delete)
		v1.GET("/product", productController.FindAll)
		v1.GET("/product/:id", productController.FindById)
	}
}
