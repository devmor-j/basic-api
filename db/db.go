package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection string

const (
	ProductsCollection Collection = "products"
)

const (
	dburi  = "mongodb://localhost:27017"
	DBname = "basic-api"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

func GetMongoClient() (*mongo.Client, error) {

	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(dburi)
		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client
		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
