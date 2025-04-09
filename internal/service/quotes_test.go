package service

import (
	"testing"

	"github.com/shawnkost/dev-quotes-api/internal/repository"
)

var mockQuotes = []repository.Quote{
	{
		ID:     "1",
		Author: "Grace Hopper",
		Text:   "The most damaging phrase in the English language is 'It's always been done this way.'",
		Tags:   []string{"culture", "leadership"},
	},
	{
		ID:     "2",
		Author: "Linus Torvalds",
		Text:   "Talk is cheap. Show me the code.",
		Tags:   []string{"programming", "open-source"},
	},
}

func TestFindQuoteById(t *testing.T) {
	id := "2"

	quote := findQuoteById(id)

	if quote == nil {
		t.Fatalf("Expected quote with ID %s, but got nil", id)
	}

	if quote.ID != id {
		t.Errorf("Expected quote with ID %s, but got %s", id, quote.ID)
	}
}

func TestFindQuoteById_NotFound(t *testing.T) {
	id := "999"

	quote := findQuoteById(id)

	if quote != nil {
		t.Errorf("Expected nil, but got quote with ID %s", id)
	}
}

func findQuoteById(id string) *repository.Quote {
	for _, quote := range mockQuotes {
		if quote.ID == id {
			return &quote
		}
	}
	return nil
}
