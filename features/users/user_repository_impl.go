package users

import (
	"context"
	"hoangcaophi/english-note-app/src/backend/global"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepositoryWriteImpl struct {
	dbCollection *mongo.Collection
}

func NewUserRepositoryWriteImpl() *UserRepositoryWriteImpl {
	return &UserRepositoryWriteImpl{
		dbCollection: global.MongoDbDatabase.Collection("users"),
	}
}

func (u *UserRepositoryWriteImpl) AddUser(user *User) (bson.Binary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := u.dbCollection.InsertOne(ctx, user)

	if err != nil {
		return bson.Binary{}, err
	}

	return result.InsertedID.(bson.Binary), nil
}

type UserRepositoryReadImpl struct {
	dbCollection *mongo.Collection
}

func NewUserRepositoryReadImpl() *UserRepositoryReadImpl {
	return &UserRepositoryReadImpl{
		dbCollection: global.MongoDbDatabase.Collection("users"),
	}
}

func (u *UserRepositoryReadImpl) GetUserByUserName(userName string) (*User, error) {
	filter := bson.D{{Key: "Username", Value: userName}}

	var userResponse *User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := u.dbCollection.FindOne(ctx, filter).Decode(&userResponse)
	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (u *UserRepositoryReadImpl) CheckUserExist(userName string) (bool, error) {
	filter := bson.D{{Key: "Username", Value: userName}}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := u.dbCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
