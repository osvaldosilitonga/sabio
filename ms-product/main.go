package main

import (
	"fmt"
	"ms-product/helpers"
	"ms-product/initializers"
	"ms-product/routes"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	initializers.LoadEnvFile()
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}

	routes.Routes(e)

	PORT := os.Getenv("PRODUCT_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", PORT)))
}
