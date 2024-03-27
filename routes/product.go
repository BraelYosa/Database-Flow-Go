package routes

import (
	"app/controllers"
	"app/helpers/connection"

	"github.com/labstack/echo/v4"
)

func Product(e *echo.Echo) {

	productGroup := e.Group("/Product", connection.CheckToken)

	productGroup.POST("/Create", controllers.CreateProduct)

	productGroup.POST("/Search", controllers.SearchProduct)

	productGroup.POST("/Update", controllers.UpdateProduct)

	productGroup.POST("/View", controllers.ViewProduct)
}
