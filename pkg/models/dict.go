package models

type DictionaryRow struct {
	ID          string   `json:"id"`
	Phrase      string   `json:"phrase"`
	Translation string   `json:"translation"`
	Examples    []string `json:"examples"`
}
