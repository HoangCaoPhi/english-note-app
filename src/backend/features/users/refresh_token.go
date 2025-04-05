package users

import "go.mongodb.org/mongo-driver/v2/bson"

type RefreshToken struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	UserID    bson.ObjectID `bson:"userId"`
	Token     string        `bson:"token"`
	CreatedAt int64         `bson:"createdAt"`
	ExpiredAt int64         `bson:"expiredAt"`
	IP        string        `bson:"ip"`
	UserAgent string        `bson:"userAgent"`
	Revoked   bool          `bson:"revoked"`
	IssuedAt  int64         `bson:"issuedAt"`
}
