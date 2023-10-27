package main

import (
	"log"
	"ncbs/api/handler"
	"ncbs/api/route"
	"ncbs/config"
	"ncbs/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// loading configurations
	appConfig, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalf("Failed to load configuration: %e", err)
	}

	awsConfig := config.GetAWSConfig()

	// initialize clients/infrastructure dependencies
	dynamodb := db.InitDynamoClient(appConfig, awsConfig)

	// initialize route handlers
	recipeHandler := handler.NewRecipeHandler(*dynamodb)

	// initialize echo server & setup routes
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	route.SetUpRoutes(e, recipeHandler)

	// start server
	e.Logger.Fatal(e.Start(":" + appConfig.Port))
}
