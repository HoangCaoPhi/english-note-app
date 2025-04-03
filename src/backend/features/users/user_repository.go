package users

import "go.mongodb.org/mongo-driver/v2/bson"

type UserRepositoryWrite interface {
	AddUser(user *User) (id bson.Binary, err error)
}

type UserRepositoryRead interface {
	GetUserByUserName(userName string) (*User, error)
	CheckUserExist(userName string) (bool, error)
}

var (
	userRepositoryRead  UserRepositoryRead
	userRepositoryWrite UserRepositoryWrite
)

func NewUserRepositoryRead() UserRepositoryRead {
	return userRepositoryRead
}

func InitUserRepositoryRead(u UserRepositoryRead) {
	userRepositoryRead = u
}

func NewUserRepositoryWrite() UserRepositoryWrite {
	return userRepositoryWrite
}

func InitUserRepositoryWrite(u UserRepositoryWrite) {
	userRepositoryWrite = u
}
