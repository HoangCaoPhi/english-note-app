package users

import (
	"context"
	"hoangcaophi/english-note-app/src/backend/global"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RefreshTokenRepositoryReadImpl struct {
	dbCollection *mongo.Collection
}

func NewRefreshTokenRepositoryReadImpl() *RefreshTokenRepositoryReadImpl {
	return &RefreshTokenRepositoryReadImpl{
		dbCollection: global.MongoDbDatabase.Collection("refresh_tokens"),
	}
}

func (r *RefreshTokenRepositoryReadImpl) GetRefreshToken(userId bson.ObjectID) (RefreshToken, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result RefreshToken
	err := r.dbCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&result)
	if err != nil {
		return RefreshToken{}, err
	}

	return result, nil
}

type RefreshTokenRepositoryWriteImpl struct {
	dbCollection *mongo.Collection
}

func NewRefreshTokenRepositoryWriteImpl() *RefreshTokenRepositoryWriteImpl {
	return &RefreshTokenRepositoryWriteImpl{
		dbCollection: global.MongoDbDatabase.Collection("refresh_tokens"),
	}
}

func (r *RefreshTokenRepositoryWriteImpl) AddRefreshToken(refreshToken RefreshToken) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.dbCollection.InsertOne(ctx, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenRepositoryWriteImpl) DeleteRefreshToken(userId bson.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.dbCollection.DeleteOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenRepositoryWriteImpl) UpdateRefreshToken(
	userId bson.ObjectID,
	refreshToken string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.dbCollection.UpdateOne(
		ctx,
		bson.M{"userId": userId, "revoked": false},
		bson.M{"$set": bson.M{"token": refreshToken, "revoked": false, "createdAt": time.Now().Unix()}},
	)

	if err != nil {
		return err
	}

	return nil
}
