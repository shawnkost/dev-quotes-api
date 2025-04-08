package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shawnkost/dev-quotes-api/internal/service"
)

// RegisterRoutes registers all v1 routes
func RegisterRoutes(g *echo.Group) {
	g.GET("/quotes/random", GetRandomQuoteHandler)
}

// GetRandomQuoteHandler godoc
// @Summary Get a random developer quote
// @Description Returns a single random quote from the dataset
// @Tags quotes
// @Produce json
// @Success 200 {object} repository.Quote
// @Failure 500 {object} map[string]string
// @Router /quotes/random [get]
func GetRandomQuoteHandler(c echo.Context) error {
	quote, err := service.GetRandomQuote()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to load quote",
		})
	}

	if quote == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No quotes found",
		})
	}

	return c.JSON(http.StatusOK, quote)
}
