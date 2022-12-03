package models

type DictionaryRow struct {
	ID          string
	Phrase      string
	Translation string
	Examples    []string
}
