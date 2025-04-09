package repository

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shawnkost/dev-quotes-api/internal/errors"
)

type Quote struct {
	ID     string   `json:"id"`
	Author string   `json:"author"`
	Text   string   `json:"text"`
	Tags   []string `json:"tags"`
}

func LoadQuotes() ([]Quote, error) {
	path := filepath.Join("configs", "quotes.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.NewInternalError("failed to read quotes file")
	}

	var quotes []Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return nil, errors.NewInternalError("failed to parse quotes data")
	}

	return quotes, nil
}

func GetQuoteByID(id string) (*Quote, error) {
	quotes, err := LoadQuotes()
	if err != nil {
		return nil, err
	}

	for _, quote := range quotes {
		if quote.ID == id {
			return &quote, nil
		}
	}

	return nil, errors.NewNotFoundError("quote not found")
}
