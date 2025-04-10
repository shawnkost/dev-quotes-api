package service

import (
	"math/rand"
	"strings"

	"github.com/shawnkost/dev-quotes-api/internal/errors"
	"github.com/shawnkost/dev-quotes-api/internal/repository"
)

func GetRandomQuote() (*repository.Quote, error) {
	quotes, err := repository.LoadQuotes()
	if err != nil {
		return nil, err
	}

	if len(quotes) == 0 {
		return nil, errors.NewNotFoundError("no quotes available")
	}

	randomIndex := rand.Intn(len(quotes))
	return &quotes[randomIndex], nil
}

func GetQuoteByID(id string) (*repository.Quote, error) {
	if id == "" {
		return nil, errors.NewValidationError("quote ID is required")
	}

	quote, err := repository.GetQuoteByID(id)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

type PaginatedQuotes struct {
	Quotes      []repository.Quote `json:"quotes"`
	Total       int                `json:"total"`
	Page        int                `json:"page"`
	PerPage     int                `json:"per_page"`
	TotalPages  int                `json:"total_pages"`
	HasNext     bool               `json:"has_next"`
	HasPrevious bool               `json:"has_previous"`
}

func GetPaginatedQuotes(author, tag string, page, perPage int) (*PaginatedQuotes, error) {
	quotes, err := repository.LoadQuotes()
	if err != nil {
		return nil, err
	}

	var filteredQuotes []repository.Quote

	for _, quote := range quotes {
		matchesAuthor := true
		matchesTag := true

		if author != "" {
			matchesAuthor = strings.Contains(strings.ToLower(quote.Author), strings.ToLower(author))
		}

		if tag != "" {
			matchesTag = false
			for _, t := range quote.Tags {
				if strings.EqualFold(t, tag) {
					matchesTag = true
					break
				}
			}
		}

		if matchesAuthor && matchesTag {
			filteredQuotes = append(filteredQuotes, quote)
		}
	}

	totalQuotes := len(filteredQuotes)
	if totalQuotes == 0 {
		return nil, errors.NewNotFoundError("no quotes found matching the provided filters")
	}

	totalPages := (totalQuotes + perPage - 1) / perPage
	startIndex := (page - 1) * perPage
	endIndex := startIndex + perPage
	if endIndex > totalQuotes {
		endIndex = totalQuotes
	}

	if startIndex >= totalQuotes {
		return nil, errors.NewValidationError("page number out of range")
	}

	result := &PaginatedQuotes{
		Quotes:      filteredQuotes[startIndex:endIndex],
		Total:       totalQuotes,
		Page:        page,
		PerPage:     perPage,
		TotalPages:  totalPages,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
	}

	return result, nil
}

func GetAllTags() ([]string, error) {
	quotes, err := repository.LoadQuotes()
	if err != nil {
		return nil, err
	}

	tagMap := make(map[string]struct{})
	for _, quote := range quotes {
		for _, tag := range quote.Tags {
			tagMap[tag] = struct{}{}
		}
	}

	tags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		tags = append(tags, tag)
	}

	return tags, nil
}
