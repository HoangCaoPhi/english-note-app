package words

import "go.mongodb.org/mongo-driver/bson/primitive"

type Word struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Word           string             `bson:"word"`
	Language       string             `bson:"language"`
	Pronunciations []Pronunciation    `bson:"pronunciations"`
	Meanings       []Meaning          `bson:"meanings"`
	Synonyms       []string           `bson:"synonyms"`
	Antonyms       []string           `bson:"antonyms"`
}
