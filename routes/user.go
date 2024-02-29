package routes

import (
	"app/controllers"

	"github.com/labstack/echo/v4"
)

func User(e *echo.Echo) {

	e.POST("User/Create", controllers.CreateUser)

	e.POST("User/Search", controllers.SearchUsers)

	e.POST("User/Update", controllers.UpdateUser)

	e.POST("User/View", controllers.ViewUser)

}
