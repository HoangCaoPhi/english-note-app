package words

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordService interface {
	CreateWord(ctx context.Context, createWordRequest CreateWordRequest) (bson.ObjectID, error)
	GetWordsByGroupID(ctx context.Context, groupID bson.ObjectID, page, limit int) ([]Word, int64, error)
}

var (
	wordService WordService
)

func InitWordService(w WordService) {
	wordService = w
}

func NewWordService() WordService {
	return wordService
}
