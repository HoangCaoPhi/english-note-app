package global

import "go.mongodb.org/mongo-driver/v2/mongo"

var (
	Config          Options
	MongoDbClient   *mongo.Client
	MongoDbDatabase *mongo.Database
)
