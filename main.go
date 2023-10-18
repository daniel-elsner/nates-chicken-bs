package main

import (
	"ncbs/api/route"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.SetUpRoutes(e)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080" // Default port if not specified
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
