package repository

import (
	"context"
	"kvasir/pkg/models"
)

type Repository interface {
	CreateDictonaryRow(context.Context, models.DictionaryRow) error
	GetDictionaryRowByID(ctx context.Context, rowID string) (models.DictionaryRow, error)
	GetDictionaryRow(ctx context.Context, phrase string) (models.DictionaryRow, error)
	// UpdateDictoryRow(context.Context, models.DictionaryRow) (*models.DictionaryRow, error)
}
