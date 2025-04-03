package users

import (
	"hoangcaophi/english-note-app/src/backend/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       bson.Binary `bson:"_id"`
	Username string      `bson:"username" json:"username"`
	Email    string      `bson:"email" json:"email"`
	Password string      `bson:"password,omitempty" json:"-"`
}

func CreateUser(userName string, email string) *User {
	return &User{
		ID: func() bson.Binary {
			id, err := utils.GenerateRandomBsonBinary()
			if err != nil {
				panic(err)
			}
			return id
		}(),
		Username: userName,
		Email:    email,
	}
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
