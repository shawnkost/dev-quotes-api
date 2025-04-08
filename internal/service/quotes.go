package service

import (
	"math/rand"

	"github.com/shawnkost/dev-quotes-api/internal/repository"
)

// GetRandomQuote returns a single random quote from the list.
func GetRandomQuote() (*repository.Quote, error) {
	quotes, err := repository.LoadQuotes()
	if err != nil {
		return nil, err
	}

	if len(quotes) == 0 {
		return nil, nil // No quotes available
	}

	randomIndex := rand.Intn(len(quotes))
	return &quotes[randomIndex], nil
}
