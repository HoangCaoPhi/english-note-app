package words

import (
	"context"
	"hoangcaophi/english-note-app/src/backend/global"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type WordRepositoryReadImpl struct {
	DbCollection mongo.Collection
}

func NewWordRepositoryReadImpl() *WordRepositoryReadImpl {
	return &WordRepositoryReadImpl{
		DbCollection: *global.MongoDbDatabase.Collection("words"),
	}
}

func (w *WordRepositoryReadImpl) IsWordExistsInGroup(ctx context.Context, groupId, userId bson.ObjectID, word string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{
		"groupId": groupId,
		"userId":  userId,
		"word":    word,
	}

	count, err := w.DbCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *WordRepositoryReadImpl) GetWordsByGroupID(ctx context.Context, groupID bson.ObjectID, page, limit int) ([]Word, int64, error) {
	filter := bson.M{"groupId": groupID}
	skip := (page - 1) * limit

	cursor, err := r.DbCollection.Find(ctx, filter, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var words []Word
	if err := cursor.All(ctx, &words); err != nil {
		return nil, 0, err
	}

	count, err := r.DbCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return words, count, nil
}

type WordRepositoryWriteImpl struct {
	DbCollection mongo.Collection
}

func NewWordRepositoryWriteImpl() *WordRepositoryWriteImpl {
	return &WordRepositoryWriteImpl{
		DbCollection: *global.MongoDbDatabase.Collection("words"),
	}
}

func (w *WordRepositoryWriteImpl) CreateWord(ctx context.Context, word Word) (bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	result, err := w.DbCollection.InsertOne(ctx, word)
	return result.InsertedID.(bson.ObjectID), err
}
