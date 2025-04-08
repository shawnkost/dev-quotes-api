package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
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
		return nil, err
	}

	var quotes []Quote
	if err := json.Unmarshal(data, &quotes); err != nil {
		return nil, err
	}

	return quotes, nil
}
