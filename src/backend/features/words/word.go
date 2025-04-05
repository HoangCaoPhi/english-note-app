package words

import "go.mongodb.org/mongo-driver/v2/bson"

type Word struct {
	ID             bson.Binary     `bson:"_id"`
	GroupID        bson.Binary     `bson:"groupId"`
	UserID         bson.Binary     `bson:"userId"`
	Word           string          `bson:"word"`
	Language       string          `bson:"language"`
	Pronunciations []Pronunciation `bson:"pronunciations"`
	Meanings       []Meaning       `bson:"meanings"`
	Synonyms       []string        `bson:"synonyms"`
	Antonyms       []string        `bson:"antonyms"`
	CreatedAt      int64           `bson:"createdAt"`
}
