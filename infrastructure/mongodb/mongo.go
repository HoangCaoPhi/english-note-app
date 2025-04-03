package mongodb

import (
	"context"
	"fmt"
	"hoangcaophi/english-note-app/src/backend/global"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitMongoDb() {
	mongoConfig := global.Config.MongoDb
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(mongoConfig.ConnectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatalf("Error when connecting to MongoDB: %v", err)
		return
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error when pinging MongoDB: %v", err)
		return
	}

	global.MongoDbClient = client
	global.MongoDbDatabase = client.Database("english-note")

	var result bson.M
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Fatalf("error when running ping command: %v", err)
		return
	}

	fmt.Println("pinged your deployment. You successfully connected to MongoDB!")
}
