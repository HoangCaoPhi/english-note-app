package users

import "go.mongodb.org/mongo-driver/v2/bson"

type (
	RefreshTokenRepositoryRead interface {
		GetRefreshToken(userId bson.Binary) (RefreshToken, error)
	}

	RefreshTokenRepositoryWrite interface {
		AddRefreshToken(refreshToken RefreshToken) error
		UpdateRefreshToken(
			userId bson.Binary,
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
