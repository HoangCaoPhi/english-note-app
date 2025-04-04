package wordgroups

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordGroup struct {
	ID        bson.Binary   `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	UserID    bson.Binary   `bson:"user_id"`
	WordIDs   []bson.Binary `bson:"word_ids"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
