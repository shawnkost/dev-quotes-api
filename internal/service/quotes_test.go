package service

import (
	"math/rand"
	"testing"

	"github.com/shawnkost/dev-quotes-api/internal/errors"
	"github.com/shawnkost/dev-quotes-api/internal/repository"
)

type MockQuoteRepository struct {
	quotes []repository.Quote
}

func (m *MockQuoteRepository) LoadQuotes() ([]repository.Quote, error) {
	return m.quotes, nil
}

func (m *MockQuoteRepository) GetQuoteByID(id string) (*repository.Quote, error) {
	for _, quote := range m.quotes {
		if quote.ID == id {
			return &quote, nil
		}
	}
	return nil, errors.NewNotFoundError("quote not found")
}

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
	mockRepo := &MockQuoteRepository{quotes: mockQuotes}
	service := NewQuoteService(mockRepo)

	id := "2"
	quote, err := service.GetQuoteByID(id)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if quote == nil {
		t.Fatalf("Expected quote with ID %s, but got nil", id)
	}

	if quote.ID != id {
		t.Errorf("Expected quote with ID %s, but got %s", id, quote.ID)
	}
}

func TestFindQuoteById_NotFound(t *testing.T) {
	mockRepo := &MockQuoteRepository{quotes: mockQuotes}
	service := NewQuoteService(mockRepo)

	id := "999"
	quote, err := service.GetQuoteByID(id)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if quote != nil {
		t.Errorf("Expected nil, but got quote with ID %s", id)
	}
}

func TestGetRandomQuote(t *testing.T) {
	mockRepo := &MockQuoteRepository{quotes: mockQuotes}
	service := NewQuoteService(mockRepo)

	quote, err := service.GetRandomQuote()

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if quote == nil {
		t.Fatal("Expected a quote, but got nil")
	}

	found := false
	for _, mockQuote := range mockQuotes {
		if quote.ID == mockQuote.ID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Quote with ID %s is not in our mock data", quote.ID)
	}
}

func TestGetQuoteByAuthor(t *testing.T) {
	mockRepo := &MockQuoteRepository{quotes: mockQuotes}
	service := NewQuoteService(mockRepo)

	quote, err := service.GetPaginatedQuotes("Linus Torvalds", "", 1, 10)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if quote == nil {
		t.Fatalf("Expected a quote, but got nil")
	}

	// verify the quote is from our mock data
	found := false
	for _, mockQuote := range mockQuotes {
		if quote.Quotes[0].ID == mockQuote.ID {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Quote with ID %s is not in our mock data", quote.Quotes[0].ID)
	}

	// verify pagination info
	if quote.Total != 1 {
		t.Errorf("Expected total of 1, but got %d", quote.Total)
	}

	if quote.Page != 1 {
		t.Errorf("Expected page of 1, but got %d", quote.Page)
	}

	if quote.PerPage != 10 {
		t.Errorf("Expected per_page of 10, but got %d", quote.PerPage)
	}

	if quote.TotalPages != 1 {
		t.Errorf("Expected total_pages of 1, but got %d", quote.TotalPages)
	}

	if quote.HasNext {
		t.Error("Expected has_next to be false")
	}

	if quote.HasPrevious {
		t.Error("Expected has_previous to be false")
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

func getRandomQuoteFromMock() (*repository.Quote, error) {
	if len(mockQuotes) == 0 {
		return nil, errors.NewNotFoundError("no quotes available")
	}

	randomIndex := rand.Intn(len(mockQuotes))
	return &mockQuotes[randomIndex], nil
}

// getPaginatedQuotesFromMock returns paginated quotes from mock data
func getPaginatedQuotesFromMock(author, tag string, page, perPage int) (*PaginatedQuotes, error) {
	var filteredQuotes []repository.Quote

	// filter by author if provided
	if author != "" {
		for _, quote := range mockQuotes {
			if quote.Author == author {
				filteredQuotes = append(filteredQuotes, quote)
			}
		}
	} else {
		filteredQuotes = mockQuotes
	}

	// filter by tag if provided
	if tag != "" {
		var tagFilteredQuotes []repository.Quote
		for _, quote := range filteredQuotes {
			for _, t := range quote.Tags {
				if t == tag {
					tagFilteredQuotes = append(tagFilteredQuotes, quote)
					break
				}
			}
		}
		filteredQuotes = tagFilteredQuotes
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
