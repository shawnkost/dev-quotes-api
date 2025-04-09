package service

import (
	"math/rand"
	"strings"

	"github.com/shawnkost/dev-quotes-api/internal/errors"
	"github.com/shawnkost/dev-quotes-api/internal/repository"
)

// GetRandomQuote returns a single random quote from the list.
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

func GetFilteredQuotes(author string, tag string) ([]repository.Quote, error) {
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

	if len(filteredQuotes) == 0 {
		return nil, errors.NewNotFoundError("no quotes found matching the provided filters")
	}

	return filteredQuotes, nil
}
