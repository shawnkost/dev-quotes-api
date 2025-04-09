package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shawnkost/dev-quotes-api/internal/errors"
	"github.com/shawnkost/dev-quotes-api/internal/service"
	"github.com/shawnkost/dev-quotes-api/internal/validation"
)

func RegisterRoutes(g *echo.Group) {
	RegisterHealthRoutes(g)

	g.GET("/quotes", GetFilteredQuotesHandler)
	g.GET("/quotes/:id", GetQuoteByIDHandler)
	g.GET("/quotes/random", GetRandomQuoteHandler)
}

// GetFilteredQuotesHandler godoc
// @Summary Get all quotes filtered by tag or author
// @Description Returns a list of quotes matching optional author and/or tag filters with pagination
// @Tags quotes
// @Produce json
// @Param author query string false "Author name (max 100 characters)"
// @Param tag query string false "Tag name (max 50 characters)"
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Success 200 {object} service.PaginatedQuotes
// @Failure 400 {object} errors.APIError
// @Failure 404 {object} errors.APIError
// @Failure 429 {object} errors.APIError "Rate limit exceeded"
// @Failure 500 {object} errors.APIError
// @Router /quotes [get]
// @Security ApiKeyAuth
func GetFilteredQuotesHandler(c echo.Context) error {
	params, err := validation.ValidateQuoteQueryParams(
		c.QueryParam("author"),
		c.QueryParam("tag"),
		c.QueryParam("page"),
		c.QueryParam("per_page"),
	)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok {
			return c.JSON(apiErr.Code, apiErr)
		}
		return c.JSON(http.StatusBadRequest, errors.NewValidationError("invalid query parameters"))
	}

	paginatedQuotes, err := service.GetPaginatedQuotes(
		params.Author,
		params.Tag,
		params.Page,
		params.PerPage,
	)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok {
			return c.JSON(apiErr.Code, apiErr)
		}
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError("failed to load quotes"))
	}

	return c.JSON(http.StatusOK, paginatedQuotes)
}

// GetQuoteByIDHandler godoc
// @Summary Get quote by ID
// @Description Retrieve a single quote by its unique ID
// @Tags quotes
// @Param id path string true "Quote ID"
// @Produce json
// @Success 200 {object} repository.Quote
// @Failure 400 {object} errors.APIError
// @Failure 404 {object} errors.APIError
// @Failure 429 {object} errors.APIError "Rate limit exceeded"
// @Failure 500 {object} errors.APIError
// @Router /quotes/{id} [get]
// @Security ApiKeyAuth
func GetQuoteByIDHandler(c echo.Context) error {
	id := c.Param("id")
	quote, err := service.GetQuoteByID(id)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok {
			return c.JSON(apiErr.Code, apiErr)
		}
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError("failed to load quote"))
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
// @Failure 404 {object} errors.APIError
// @Failure 429 {object} errors.APIError "Rate limit exceeded"
// @Failure 500 {object} errors.APIError
// @Router /quotes/random [get]
// @Security ApiKeyAuth
func GetRandomQuoteHandler(c echo.Context) error {
	quote, err := service.GetRandomQuote()
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok {
			return c.JSON(apiErr.Code, apiErr)
		}
		return c.JSON(http.StatusInternalServerError, errors.NewInternalError("failed to load quote"))
	}

	if quote == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "No quotes found",
		})
	}

	return c.JSON(http.StatusOK, quote)
}
