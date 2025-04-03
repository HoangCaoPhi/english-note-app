package utils

import (
	"crypto/rand"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GenerateRandomBsonBinary() (bson.Binary, error) {
	randomData := make([]byte, 16)
	_, err := rand.Read(randomData)
	if err != nil {
		return bson.Binary{}, err
	}

	return bson.Binary{
		Subtype: bson.TypeBinaryUUID,
		Data:    randomData,
	}, nil
}
