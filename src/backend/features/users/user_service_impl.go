package users

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserServiceImpl struct {
	userRepositoryRead  UserRepositoryRead
	userRepositoryWrite UserRepositoryWrite
}

func NewUserServiceImpl(userRepositoryRead UserRepositoryRead,
	userRepositoryWrite UserRepositoryWrite) *UserServiceImpl {
	return &UserServiceImpl{
		userRepositoryRead:  userRepositoryRead,
		userRepositoryWrite: userRepositoryWrite,
	}
}

func (u *UserServiceImpl) Login(context context.Context, user *User) {

}

func (u *UserServiceImpl) Register(ctx context.Context, user *User) (bson.Binary, error) {
	userExist, err := u.userRepositoryRead.CheckUserExist(user.Username)
	if err != nil {
		log.Printf("Error when checking if user exists: %v", err)
		return bson.Binary{}, errors.New("error checking if user exists")
	}

	if userExist {
		return bson.Binary{}, errors.New("username or Email already exists")
	}

	userID, err := u.userRepositoryWrite.AddUser(user)
	if err != nil {
		log.Printf("Error when adding user: %v", err)
		return bson.Binary{}, errors.New("error registering user")
	}

	return userID, nil
}
