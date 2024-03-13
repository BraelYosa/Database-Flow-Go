package main

import (
	"app/helpers/connection"
	"app/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	data, err := connection.InitDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to Database")
	}

	connection.DB = data

	routes.Admin(e)

	routes.User(e)

	e.Logger.Fatal(e.Start(":1000"))

}
