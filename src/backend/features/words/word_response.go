package words

type PronunciationResponse struct {
	IPA    string `json:"ipa"`
	Region string `json:"region"`
}

type MeaningResponse struct {
	PartOfSpeech string   `json:"partOfSpeech"`
	Definition   string   `json:"definition"`
	Examples     []string `json:"examples"`
}

type WordResponse struct {
	ID             string                  `json:"id"`
	GroupID        string                  `json:"groupId"`
	Word           string                  `json:"word"`
	Language       string                  `json:"language"`
	Pronunciations []PronunciationResponse `json:"pronunciations"`
	Meanings       []MeaningResponse       `json:"meanings"`
	Synonyms       []string                `json:"synonyms"`
	Antonyms       []string                `json:"antonyms"`
	CreatedAt      int64                   `json:"createdAt"`
}

func (w *Word) ToResponse() WordResponse {
	return WordResponse{
		ID:             w.ID.Hex(),
		GroupID:        w.GroupID.Hex(),
		Word:           w.Word,
		Language:       w.Language,
		Pronunciations: toPronunciationResponses(w.Pronunciations),
		Meanings:       toMeaningResponses(w.Meanings),
		Synonyms:       w.Synonyms,
		Antonyms:       w.Antonyms,
		CreatedAt:      w.CreatedAt,
	}
}

func toPronunciationResponses(pronunciations []Pronunciation) []PronunciationResponse {
	responses := make([]PronunciationResponse, len(pronunciations))
	for i, p := range pronunciations {
		responses[i] = PronunciationResponse{
			IPA:    p.IPA,
			Region: p.Region,
		}
	}
	return responses
}

func toMeaningResponses(meanings []Meaning) []MeaningResponse {
	responses := make([]MeaningResponse, len(meanings))
	for i, m := range meanings {
		responses[i] = MeaningResponse{
			PartOfSpeech: m.PartOfSpeech,
			Definition:   m.Definition,
			Examples:     m.Examples,
		}
	}
	return responses
}

type WordListResponse struct {
	Data       []WordResponse     `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}
