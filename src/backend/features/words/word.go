package words

import "go.mongodb.org/mongo-driver/v2/bson"

type Word struct {
	ID             bson.ObjectID   `bson:"_id"`
	GroupID        bson.ObjectID   `bson:"groupId"`
	UserID         bson.ObjectID   `bson:"userId"`
	Word           string          `bson:"word"`
	Language       string          `bson:"language"`
	Pronunciations []Pronunciation `bson:"pronunciations"`
	Meanings       []Meaning       `bson:"meanings"`
	Synonyms       []string        `bson:"synonyms"`
	Antonyms       []string        `bson:"antonyms"`
	CreatedAt      int64           `bson:"createdAt"`
}
