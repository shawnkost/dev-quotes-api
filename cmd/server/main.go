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

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
		MaxAge: 300,
	}))

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            31536000,
		HSTSExcludeSubdomains: false,
	}))

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Quote API Routes
	api := e.Group("/v1")
	v1.RegisterRoutes(api)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
