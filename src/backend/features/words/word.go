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

func CreateWord(
	groupId bson.ObjectID,
	userId bson.ObjectID,
	word string,
	language string,
	pronunciations []Pronunciation,
	meanings []Meaning,
	synonyms []string,
	antonyms []string,
) *Word {
	return &Word{
		ID:             bson.NewObjectID(),
		GroupID:        groupId,
		UserID:         userId,
		Word:           word,
		Language:       language,
		Pronunciations: pronunciations,
		Meanings:       meanings,
		Synonyms:       synonyms,
		Antonyms:       antonyms,
	}
}
