package users

import (
	"context"
	"errors"
	"hoangcaophi/english-note-app/src/backend/global"
	"log"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserServiceImpl struct {
	userRepositoryRead          UserRepositoryRead
	userRepositoryWrite         UserRepositoryWrite
	refreshTokenRepositoryRead  RefreshTokenRepositoryRead
	refreshTokenRepositoryWrite RefreshTokenRepositoryWrite
}

func NewUserServiceImpl(
	userRepositoryRead UserRepositoryRead,
	userRepositoryWrite UserRepositoryWrite,
	refreshTokenRead RefreshTokenRepositoryRead,
	refreshTokenWrite RefreshTokenRepositoryWrite) *UserServiceImpl {
	return &UserServiceImpl{
		userRepositoryRead:          userRepositoryRead,
		userRepositoryWrite:         userRepositoryWrite,
		refreshTokenRepositoryRead:  refreshTokenRead,
		refreshTokenRepositoryWrite: refreshTokenWrite,
	}
}

func (u *UserServiceImpl) Register(ctx context.Context, user *User) (bson.ObjectID, error) {
	userExist, err := u.userRepositoryRead.CheckUserExist(user.Username)
	if err != nil {
		log.Printf("Error when checking if user exists: %v", err)
		return bson.ObjectID{}, errors.New("error checking if user exists")
	}

	if userExist {
		return bson.ObjectID{}, errors.New("username or Email already exists")
	}

	userID, err := u.userRepositoryWrite.AddUser(user)
	if err != nil {
		log.Printf("Error when adding user: %v", err)
		return bson.ObjectID{}, errors.New("error registering user")
	}

	return userID, nil
}

func (u *UserServiceImpl) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := userRepositoryRead.GetUserByUserName(username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", errors.New("invalid username or password")
		}
		log.Printf("Error when querying user: %v", err)
		return "", "", errors.New("internal server error")
	}

	if !user.VerifyPassword(password) {
		return "", "", errors.New("invalid username or password")
	}

	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.generateRefreshToken(ctx, user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func getIP(ctx context.Context) string {
	ip, ok := ctx.Value("IP").(string)
	if !ok {
		return ""
	}
	return ip
}

func getUserAgent(ctx context.Context) string {
	userAgent, ok := ctx.Value("UserAgent").(string)
	if !ok {
		return ""
	}
	return userAgent
}

func generateAccessToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"sub":      user.ID.Hex(),
		"exp":      time.Now().Add(time.Minute * 15).Unix(),
		"username": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(global.Config.Authentication.Jwt.SecretKey)
	accessToken, err := token.SignedString(secretKey)

	if err != nil {
		log.Printf("Error signing access token: %v", err)
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}

func (u *UserServiceImpl) generateRefreshToken(ctx context.Context, user *User) (string, error) {
	token := generateRandomString(32)

	err := u.saveRefreshToken(ctx, user.ID, token)
	if err != nil {
		log.Printf("Error saving refresh token: %v", err)
		return "", errors.New("failed to generate refresh token")
	}

	return token, nil
}

func (u *UserServiceImpl) saveRefreshToken(ctx context.Context, userID bson.ObjectID, refreshToken string) error {
	expiredAt := time.Now().Add(time.Hour * 24 * 7).Unix()

	ip := getIP(ctx)
	userAgent := getUserAgent(ctx)

	refreshTokenRecord := RefreshToken{
		UserID:    userID,
		Token:     refreshToken,
		CreatedAt: time.Now().Unix(),
		ExpiredAt: expiredAt,
		IP:        ip,
		UserAgent: userAgent,
		Revoked:   false,
		IssuedAt:  time.Now().Unix(),
	}

	err := u.refreshTokenRepositoryWrite.AddRefreshToken(refreshTokenRecord)
	if err != nil {
		log.Printf("Error saving refresh token: %v", err)
		return errors.New("failed to save refresh token")
	}

	return nil
}

func generateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, letters[rand.Intn(len(letters))])
	}
	return string(result)
}

func (u *UserServiceImpl) RefreshAccessToken(ctx context.Context, refreshToken string, userId bson.ObjectID) (string, error) {
	valid, err := u.isRefreshTokenValid(userId, refreshToken, ctx)
	if err != nil || !valid {
		return "", errors.New("invalid refresh token")
	}

	refreshTokenRecord, err := u.refreshTokenRepositoryRead.GetRefreshToken(userId)
	if err != nil {
		return "", err
	}

	userID := refreshTokenRecord.UserID

	user, err := u.userRepositoryRead.GetByUserId(userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("user not found")
		}
		log.Printf("Error when querying user: %v", err)
		return "", errors.New("internal server error")
	}

	accessToken, err := generateAccessToken(user)
	if err != nil {
		return "", err
	}

	newRefreshToken, err := u.generateRefreshToken(ctx, user)
	if err != nil {
		return "", err
	}

	ip := ctx.Value("IP").(string)
	userAgent := ctx.Value("UserAgent").(string)

	err = u.updateRefreshToken(user.ID, newRefreshToken, ip, userAgent)
	if err != nil {
		return "", errors.New("failed to update refresh token")
	}

	return accessToken, nil
}

func (u *UserServiceImpl) updateRefreshToken(
	userID bson.ObjectID,
	newRefreshToken string,
	id string,
	userAgent string) error {

	err := refreshTokenRepositoryWrite.UpdateRefreshToken(userID, newRefreshToken)

	if err != nil {
		return errors.New("failed to update refresh token")
	}

	return nil
}

func (u *UserServiceImpl) isRefreshTokenValid(userId bson.ObjectID, refreshToken string, ctx context.Context) (bool, error) {
	tokenRecord, err := refreshTokenRepositoryRead.GetRefreshToken(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, errors.New("refresh token not found")
		}
		return false, errors.New("error checking refresh token")
	}

	if tokenRecord.ExpiredAt < time.Now().Unix() {
		return false, errors.New("refresh token expired")
	}

	if tokenRecord.Revoked {
		return false, errors.New("refresh token revoked")
	}

	ip := ctx.Value("IP").(string)
	userAgent := ctx.Value("UserAgent").(string)
	if tokenRecord.IP != ip || tokenRecord.UserAgent != userAgent {
		return false, errors.New("invalid refresh token usage")
	}

	return true, nil
}
