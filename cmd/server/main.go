package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/shawnkost/dev-quotes-api/docs" // Swagger docs
	v1 "github.com/shawnkost/dev-quotes-api/internal/api/v1"
	"github.com/shawnkost/dev-quotes-api/internal/config"
	"github.com/shawnkost/dev-quotes-api/internal/errors"
	"github.com/shawnkost/dev-quotes-api/internal/logger"
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
	log := logger.Logger()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load configuration")
		return
	}

	e := echo.New()

	// custom logger middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			latency := time.Since(start)

			req := c.Request()
			res := c.Response()

			log.Info().
				Str("method", req.Method).
				Str("uri", req.RequestURI).
				Int("status", res.Status).
				Dur("latency", latency).
				Str("ip", c.RealIP()).
				Str("user_agent", req.UserAgent()).
				Msg("request completed")

			return err
		}
	})

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

		log.Error().
			Err(err).
			Int("status", code).
			Str("method", c.Request().Method).
			Str("uri", c.Request().RequestURI).
			Str("ip", c.RealIP()).
			Msg("request failed")

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

	log.Info().
		Str("port", cfg.Server.Port).
		Str("environment", cfg.Server.Environment).
		Msg("starting server")

	if err := e.StartServer(s); err != nil {
		log.Fatal().Err(err).Msg("server failed to start")
	}
}
