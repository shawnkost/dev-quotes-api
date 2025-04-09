package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/shawnkost/dev-quotes-api/docs" // Swagger docs
	v1 "github.com/shawnkost/dev-quotes-api/internal/api/v1"
	"github.com/shawnkost/dev-quotes-api/internal/config"
	"github.com/shawnkost/dev-quotes-api/internal/errors"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Dev Quotes API
// @version 1.0
// @description Public API for developer-related quotes; Rate limiting is applied per IP address. Default rate limit is 50 requests per minute. Random quote endpoint has a higher limit of 100 requests per minute.
// @contact.name Shawn Kost
// @contact.url https://github.com/shawnkost
// @contact.email shawnmkost@gmail.com
// @host localhost:8080
// @BasePath /v1
func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		return
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	rateLimiter := middleware.NewRateLimiterMemoryStore(50)
	e.Use(middleware.RateLimiter(rateLimiter))

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

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code    = http.StatusInternalServerError
			message interface{}
		)

		if apiErr, ok := err.(*errors.APIError); ok {
			code = apiErr.Code
			message = apiErr
		} else if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message
		} else {
			message = errors.NewInternalError("internal server error")
		}

		c.Logger().Error(err)

		if !c.Response().Committed {
			c.JSON(code, message)
		}
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/v1")
	v1.RegisterRoutes(api)

	s := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	e.Logger.Fatal(e.StartServer(s))
}
