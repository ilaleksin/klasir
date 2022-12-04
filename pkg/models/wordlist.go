package models

type Review struct {
	ID           string
	WordsChecked []int
	Words        []DictionaryRow
	NumCorrect   int
	NumMistakes  int
}

type WordList struct {
	ID    string
	Words []DictionaryRow
}

type ReviewOptions struct {
	Length int    `json:"length"`
	Filter Filter `json:"filter"`
}

type Filter struct {
	Topic string `json:"topic"`
}
