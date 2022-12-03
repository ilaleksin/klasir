package repository

import (
	"context"
	"errors"
	"kvasir/pkg/models"
)

type InMemoryStorage struct {
	dict map[string]models.DictionaryRow
}

func NewInMemoryStorage() *InMemoryStorage {
	memDict := make(map[string]models.DictionaryRow)
	return &InMemoryStorage{
		dict: memDict,
	}
}

func (m *InMemoryStorage) CreateDictonaryRow(ctx context.Context, row models.DictionaryRow) error {
	if _, ok := m.dict[row.Phrase]; ok {
		return errors.New("phrase already exists")
	}
	m.dict[row.Phrase] = row
	return nil
}

func (m *InMemoryStorage) GetPhrase(ctx context.Context)
