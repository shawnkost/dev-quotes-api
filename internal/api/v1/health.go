package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func RegisterHealthRoutes(g *echo.Group) {
	g.GET("/health", HealthCheckHandler)
}

// HealthCheckHandler godoc
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthResponse{
		Status:  "ok",
		Version: "1.0.0",
	})
}
