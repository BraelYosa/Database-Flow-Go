package routes

import (
	"app/controllers"
	"app/helpers/connection"

	"github.com/labstack/echo/v4"
)

func Admin(e *echo.Echo) {

	e.POST("Signup", controllers.Signup)

	e.POST("Login", controllers.Login)

}

func User(e *echo.Echo) {

	userGroup := e.Group("/User", connection.CheckToken)

	userGroup.POST("/Create", controllers.CreateUser)

	userGroup.POST("/Search", controllers.SearchUsers)

	userGroup.POST("/Update", controllers.UpdateUser)

	userGroup.POST("/View", controllers.ViewUser)
}
