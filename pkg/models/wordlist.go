package models

type WordList struct {
	Words []DictionaryRow
}

type Review struct {
	WordsChecked []int
	Words        []DictionaryRow
	Correct      int
	Mistake      int
}
