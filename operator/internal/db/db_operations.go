package db

import (
	"context"
	"os"
	"time"

	"github.com/dana-team/axiom/operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertClusterInfoToMongo stores or updates cluster information in MongoDB database.
// It connects to MongoDB using the MONGO_URI environment variable, and performs an upsert operation
// based on the cluster ID. The function handles the connection lifecycle, including proper cleanup
// and error logging. If the cluster ID is empty, no operation is performed.
func InsertClusterInfoToMongo(logger logr.Logger, clusterInfo v1alpha1.ClusterInfo) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		logger.Error(nil, "MONGO_URI environment variable not set")
		return
	}
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		logger.Error(err, "Failed to connect to MongoDB")
		return
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			return
		}
	}(client, ctx)

	if clusterInfo.Status.ClusterID != "" {
		collection := client.Database("axiom").Collection("clusterInfo")
		filter := bson.M{"clusterID": clusterInfo.Status.ClusterID}
		update := bson.M{"$set": clusterInfo.Status}
		opts := options.Update().SetUpsert(true)
		_, err = collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			logger.Error(err, "Failed to insert cluster info to MongoDB")
			return
		}
		logger.Info("Inserted cluster info to MongoDB")
	}
}
