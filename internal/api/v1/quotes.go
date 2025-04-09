package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shawnkost/dev-quotes-api/internal/service"
)

// RegisterRoutes registers all v1 routes
func RegisterRoutes(g *echo.Group) {
	g.GET("/quotes", GetFilteredQuotesHandler)
	g.GET("/quotes/:id", GetQuoteByIDHandler)
	g.GET("/quotes/random", GetRandomQuoteHandler)
}

// GetFilteredQuotesHandler godoc
// @Summary Get all quotes filtered by tag or author
// @Description Returns a list of quotes matching optional author and/or tag filters
// @Tags quotes
// @Produce json
// @Param author query string false "Author name"
// @Param tag query string false "Tag name"
// @Success 200 {array} repository.Quote
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /quotes [get]
func GetFilteredQuotesHandler(c echo.Context) error {
	author := c.QueryParam("author")
	tag := c.QueryParam("tag")

	quotes, err := service.GetFilteredQuotes(author, tag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to load quotes",
		})
	}

	if len(quotes) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No quotes found matching the provided filters",
		})
	}

	return c.JSON(http.StatusOK, quotes)
}

// GetQuoteByIDHandler godoc
// @Summary Get quote by ID
// @Description Retrieve a single quote by its unique ID
// @Tags quotes
// @Param id path string true "Quote ID"
// @Produce json
// @Success 200 {object} repository.Quote
// @Failure 404 {object} map[string]string
// @Router /quotes/{id} [get]
func GetQuoteByIDHandler(c echo.Context) error {
	id := c.Param("id")
	quote, err := service.GetQuoteByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to load quote",
		})
	}

	if quote == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Quote not found",
		})
	}

	return c.JSON(http.StatusOK, quote)

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
