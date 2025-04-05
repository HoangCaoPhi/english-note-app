package wordgroups

import (
	"hoangcaophi/english-note-app/src/backend/shared"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordGroup struct {
	ID        bson.Binary `bson:"_id,omitempty"`
	UserID    bson.Binary `bson:"userId"`
	Name      string      `bson:"name"`
	CreatedAt int64       `bson:"createdAt"`
}

func NewWordGroup(userId bson.Binary, name string) *WordGroup {
	return &WordGroup{
		ID: func() bson.Binary {
			id, err := shared.GenerateRandomBsonBinary()
			if err != nil {
				panic(err)
			}
			return id
		}(),
		UserID:    userId,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	}
}
