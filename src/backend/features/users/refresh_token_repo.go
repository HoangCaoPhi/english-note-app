package users

import "go.mongodb.org/mongo-driver/v2/bson"

type (
	RefreshTokenRepositoryRead interface {
		GetRefreshToken(userId bson.ObjectID) (RefreshToken, error)
	}

	RefreshTokenRepositoryWrite interface {
		AddRefreshToken(refreshToken RefreshToken) error
		UpdateRefreshToken(
			userId bson.ObjectID,
			refreshToken string) error
	}
)

var (
	refreshTokenRepositoryRead  RefreshTokenRepositoryRead
	refreshTokenRepositoryWrite RefreshTokenRepositoryWrite
)

func NewRefreshTokenRepositoryRead() RefreshTokenRepositoryRead {
	return refreshTokenRepositoryRead
}

func InitRefreshTokenRepositoryRead(read RefreshTokenRepositoryRead) {
	refreshTokenRepositoryRead = read
}

func NewRefreshTokenRepositoryWrite() RefreshTokenRepositoryWrite {
	return refreshTokenRepositoryWrite
}

func InitRefreshTokenRepositoryWrite(write RefreshTokenRepositoryWrite) {
	refreshTokenRepositoryWrite = write
}
