package validation

import (
	"strconv"

	"github.com/shawnkost/dev-quotes-api/internal/errors"
)

type QuoteQueryParams struct {
	Author  string
	Tag     string
	Page    int
	PerPage int
}

func ValidateQuoteQueryParams(author, tag, pageStr, perPageStr string) (*QuoteQueryParams, error) {
	params := &QuoteQueryParams{
		Author:  author,
		Tag:     tag,
		Page:    1,  // default page
		PerPage: 10, // default per page
	}

	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return nil, errors.NewValidationError("page must be a positive integer")
		}
		params.Page = page
	}

	if perPageStr != "" {
		perPage, err := strconv.Atoi(perPageStr)
		if err != nil || perPage < 1 {
			return nil, errors.NewValidationError("per_page must be a positive integer")
		}
		if perPage > 100 {
			return nil, errors.NewValidationError("per_page cannot exceed 100")
		}
		params.PerPage = perPage
	}

	if len(author) > 100 {
		return nil, errors.NewValidationError("author name too long (max 100 characters)")
	}

	if len(tag) > 50 {
		return nil, errors.NewValidationError("tag name too long (max 50 characters)")
	}

	return params, nil
}
