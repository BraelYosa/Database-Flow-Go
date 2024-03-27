package main

import (
	"app/helpers/connection/postgres"
	"app/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	data, err := postgres.InitDB()
	if err != nil {
		e.Logger.Fatal("Failed to connect to Database")
	}

	postgres.DB = data

	routes.Admin(e)

	routes.User(e)

	routes.Product(e)

	e.Logger.Fatal(e.Start(":1000"))

}
