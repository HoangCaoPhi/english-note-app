package users

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserService interface {
	Register(ctx context.Context, user *User) (bson.Binary, error)
	Login(ctx context.Context, user *User)
}

var (
	userService UserService
)

func NewUserService() UserService {
	if userService == nil {
		panic("No implementation was found for UserService")
	}

	return userService
}

func InitUserService(userSevice UserService) {
	userService = userSevice
}
