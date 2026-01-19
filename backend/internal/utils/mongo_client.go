package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// InitMongoDB initializes a new MongoDB connection using environment variables
// for connection string and database name. It establishes and verifies the connection
// by performing a ping test
func InitMongoDB() (*MongoClient, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return nil, fmt.Errorf("DB_NAME is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return &MongoClient{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

// Disconnect gracefully closes the MongoDB connection, ensuring all resources
// are properly released
func (mc *MongoClient) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mc.Client.Disconnect(ctx)
}

// GetCollection returns a MongoDB collection instance for the specified collection name
// from the configured database
func (mc *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return mc.Database.Collection(collectionName)
}
