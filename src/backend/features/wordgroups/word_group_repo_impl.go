package wordgroups

import (
	"context"
	"hoangcaophi/english-note-app/src/backend/global"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// READ REPOSITORY
type WordGroupRepositoryReadImpl struct {
	dbCollection *mongo.Collection
}

func NewWordGroupRepositoryReadImpl() *WordGroupRepositoryReadImpl {
	return &WordGroupRepositoryReadImpl{
		dbCollection: global.MongoDbDatabase.Collection("word_groups"),
	}
}

func (wordgroup *WordGroupRepositoryReadImpl) GetWordGroupsByUserID(userID bson.Binary) ([]WordGroup, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := wordgroup.dbCollection.Find(ctx, map[string]any{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var wordGroups []WordGroup
	for cursor.Next(ctx) {
		var wordGroup WordGroup
		if err := cursor.Decode(&wordGroup); err != nil {
			return nil, err
		}
		wordGroups = append(wordGroups, wordGroup)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return wordGroups, nil
}

// WRITE REPOSITORY
type WordGroupRepositoryWriteImpl struct {
	dbCollection *mongo.Collection
}

func NewWordGroupRepositoryWriteImpl() *WordGroupRepositoryWriteImpl {
	return &WordGroupRepositoryWriteImpl{
		dbCollection: global.MongoDbDatabase.Collection("word_groups"),
	}
}

func (wordgroup *WordGroupRepositoryWriteImpl) CreateWordGroup(wordGroup *WordGroup) (bson.Binary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, err := wordgroup.dbCollection.InsertOne(ctx, wordGroup)
	if err != nil {
		return bson.Binary{}, err
	}
	return id.InsertedID.(bson.Binary), nil
}

func (wordgroup *WordGroupRepositoryWriteImpl) UpdateWordGroup(wordGroup WordGroup) (bson.Binary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := wordgroup.dbCollection.UpdateOne(ctx, map[string]any{"_id": wordGroup.ID}, map[string]any{"$set": wordGroup})
	if err != nil {
		return bson.Binary{}, err
	}
	return wordGroup.ID, nil
}

func (wordgroup *WordGroupRepositoryWriteImpl) DeleteWordGroup(id bson.Binary) (bson.Binary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := wordgroup.dbCollection.DeleteOne(ctx, map[string]any{"_id": id})
	if err != nil {
		return bson.Binary{}, err
	}
	return id, nil
}
