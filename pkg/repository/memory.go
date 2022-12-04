package repository

import (
	"context"
	"errors"
	"fmt"
	"kvasir/pkg/models"
	"math/rand"
	"sort"
	"time"
)

type InMemoryStorage struct {
	sortedKeys   []string
	dictByPhrase map[string]models.DictionaryRow
	dictByID     map[string]models.DictionaryRow
}

func NewInMemoryStorage() *InMemoryStorage {
	dictPhrase := make(map[string]models.DictionaryRow)
	dictID := make(map[string]models.DictionaryRow)
	sortedKeys := []string{}
	return &InMemoryStorage{
		dictByPhrase: dictPhrase,
		dictByID:     dictID,
		sortedKeys:   sortedKeys,
	}
}

func (m *InMemoryStorage) CreateDictonaryRow(ctx context.Context, row models.DictionaryRow) error {
	if _, ok := m.dictByPhrase[row.Phrase]; ok {
		return errors.New("phrase already exists")
	}
	m.dictByPhrase[row.Phrase] = row
	m.dictByID[row.ID] = row
	m.insertPhraseSorted(row.Phrase)
	return nil
}

func (m *InMemoryStorage) GetDictionaryRow(ctx context.Context, phrase string) (models.DictionaryRow, error) {
	if res, ok := m.dictByPhrase[phrase]; ok {
		return res, nil
	}
	err := fmt.Errorf("phrase %s not found", phrase)
	return models.DictionaryRow{}, err
}

func (m *InMemoryStorage) GetDictionary(ctx context.Context, limit, offset int) (res []models.DictionaryRow, err error) {
	var start, end int
	if offset >= len(m.sortedKeys) {
		err = fmt.Errorf("offset %d is out of range", offset)
		return res, err
	}
	start = offset
	if start+limit > len(m.sortedKeys) {
		end = len(m.sortedKeys)
	} else {
		end = start + limit
	}
	for _, phrase := range m.sortedKeys[start:end] {
		res = append(res, m.dictByPhrase[phrase])
	}
	return res, err
}

func (m *InMemoryStorage) MakeWordList(ctx context.Context, options models.ReviewOptions) (models.WordList, error) {
	uniqueIndices := make(map[int]interface{})
	words := []models.DictionaryRow{}
	rand.Seed(time.Now().Unix())
	length := options.Length
	if length >= len(m.sortedKeys) {
		length = len(m.sortedKeys)
	}
	ind := rand.Intn(len(m.sortedKeys))
	for i := 0; i < length; i++ {
		for _, ok := uniqueIndices[ind]; ok; {
			ind = rand.Intn(len(m.sortedKeys))
		}
	}
}

func (m *InMemoryStorage) insertPhraseSorted(newPhrase string) {
	i := sort.SearchStrings(m.sortedKeys, newPhrase)
	m.sortedKeys = append(m.sortedKeys, "")
	copy(m.sortedKeys[i+1:], m.sortedKeys[i:])
	m.sortedKeys[i] = newPhrase
}
