package repository

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shawnkost/dev-quotes-api/internal/errors"
)

type QuoteRepository interface {
	LoadQuotes() ([]Quote, error)
	GetQuoteByID(id string) (*Quote, error)
}

type Quote struct {
	ID     string   `json:"id"`
	Author string   `json:"author"`
	Text   string   `json:"text"`
	Tags   []string `json:"tags"`
}

type FileQuoteRepository struct {
	filePath string
}

func NewFileQuoteRepository() *FileQuoteRepository {
	return &FileQuoteRepository{
		filePath: filepath.Join("configs", "quotes.json"),
	}
}

func (r *FileQuoteRepository) LoadQuotes() ([]Quote, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, errors.NewInternalError("failed to read quotes file")
	}

	var quotes []Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return nil, errors.NewInternalError("failed to parse quotes data")
	}

	return quotes, nil
}

func (r *FileQuoteRepository) GetQuoteByID(id string) (*Quote, error) {
	quotes, err := r.LoadQuotes()
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
