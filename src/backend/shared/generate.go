package shared

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

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

func BsonBinaryToString(bin bson.Binary) string {
	return base64.StdEncoding.EncodeToString(bin.Data)
}

func StringToBsonBinary(s string) (bson.Binary, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return bson.Binary{}, errors.New("invalid base64 string")
	}
	return bson.Binary{
		Subtype: bson.TypeBinaryUUID,
		Data:    data,
	}, nil
}
