package testutils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectTestMongo() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("TEST_MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}
