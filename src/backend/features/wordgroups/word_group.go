package wordgroups

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordGroup struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    bson.ObjectID `bson:"user_id" json:"userId"`
	Name      string        `bson:"name" json:"name"`
	CreatedAt int64         `bson:"created_at" json:"createdAt"`
}

func NewWordGroup(userId bson.ObjectID, name string) *WordGroup {
	return &WordGroup{
		ID:        bson.NewObjectID(),
		UserID:    userId,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
}
