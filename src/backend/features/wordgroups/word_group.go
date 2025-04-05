package wordgroups

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordGroup struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	UserID    bson.ObjectID `bson:"userId"`
	Name      string        `bson:"name"`
	CreatedAt int64         `bson:"createdAt"`
}

func NewWordGroup(userId bson.ObjectID, name string) *WordGroup {
	return &WordGroup{
		ID:        bson.NewObjectID(),
		UserID:    userId,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
}
