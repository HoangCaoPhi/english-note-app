package words

import "go.mongodb.org/mongo-driver/v2/bson"

type CreateWordRequest struct {
	GroupID        bson.ObjectID   `json:"groupId"`
	Word           string          `json:"word"`
	Language       string          `json:"language"`
	Pronunciations []Pronunciation `json:"pronunciations"`
	Meanings       []Meaning       `json:"meanings"`
	Synonyms       []string        `json:"synonyms"`
	Antonyms       []string        `json:"antonyms"`
}
