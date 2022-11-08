package dictionary

type Response struct {
	Word     string    `json:"word"`
	Meanings []meaning `json:"meanings"`
}

type meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []definition `json:"definitions"`
}

type definition struct {
	Definition string   `json:"definition"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}
