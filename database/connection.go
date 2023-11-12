package database

import (
	"context"
	"fmt"
	"github.com/gofor-little/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

var clientOptions *options.ClientOptions

var client *mongo.Client

func init() {
	if err := env.Load(".env"); err != nil {
		panic(err)
	}

	host := env.Get("MONGO_HOST", "")
	port, _ := strconv.Atoi(env.Get("MONGO_PORT", "0"))
	user := env.Get("MONGO_USER", "")
	password := env.Get("MONGO_PASSWORD", "")

	connectionUri := fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port)

	clientOptions = options.Client().ApplyURI(connectionUri)

	var connectionError error

	client, connectionError = mongo.Connect(context.Background(), clientOptions)

	if connectionError != nil {
		panic(connectionError)
	}
}

func GetMongoCollectionCollection(databaseName string, collectionName string) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}
