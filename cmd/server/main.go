package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/shawnkost/dev-quotes-api/docs" // Swagger docs
	v1 "github.com/shawnkost/dev-quotes-api/internal/api/v1"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Dev Quotes API
// @version 1.0
// @description Public API for developer-related quotes
// @contact.name Shawn Kost
// @contact.url https://github.com/shawnkost
// @contact.email shawnmkost@gmail.com
// @host localhost:8080
// @BasePath /v1
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Quote API Routes
	api := e.Group("/v1")
	v1.RegisterRoutes(api)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
