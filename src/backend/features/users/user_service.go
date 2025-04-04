package users

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserService interface {
	Register(ctx context.Context, user *User) (bson.Binary, error)
	Login(ctx context.Context, username, password string) (string, string, error)
	RefreshAccessToken(ctx context.Context, refreshToken string, userId bson.Binary) (string, error)
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
