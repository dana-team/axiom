package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dana-team/axiom/operator/api/v1alpha1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Default configuration
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "axiom"
	}

	fmt.Printf("Connecting to MongoDB at %s, database: %s\n", mongoURI, dbName)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect: %v", err)
		}
	}()

	db := client.Database(dbName)
	collection := db.Collection("clusterInfo")

	// Clean up existing data
	fmt.Println("Cleaning up existing data...")
	if _, err := collection.DeleteMany(ctx, bson.M{}); err != nil {
		log.Fatalf("Failed to clean up collection: %v", err)
	}

	// Create placeholder data
	clusters := []interface{}{
		v1alpha1.ClusterInfoStatus{
			Name:              "cluster-alpha",
			ClusterID:         "cid-alpha-001",
			KubernetesVersion: "v1.29.0",
			ClusterResources: v1alpha1.ClusterResources{
				CPU:    "100",
				Memory: "200",
			},
			NodeInfo: []v1alpha1.NodeInfo{
				{Name: "node-1"},
				{Name: "node-2"},
			},
			ApiServerAddresses: []string{"https://api.alpha.example.com"},
		},
		v1alpha1.ClusterInfoStatus{
			Name:              "cluster-beta",
			ClusterID:         "cid-beta-002",
			KubernetesVersion: "v1.28.5",
			ClusterResources: v1alpha1.ClusterResources{
				CPU:    "50",
				Memory: "100",
			},
			NodeInfo: []v1alpha1.NodeInfo{
				{Name: "node-1"},
			},
			ApiServerAddresses: []string{"https://api.beta.example.com"},
		},
		v1alpha1.ClusterInfoStatus{
			Name:              "cluster-gamma",
			ClusterID:         "cid-gamma-003",
			KubernetesVersion: "v1.30.0",
			ClusterResources: v1alpha1.ClusterResources{
				CPU:    "200",
				Memory: "400",
			},
			NodeInfo: []v1alpha1.NodeInfo{
				{Name: "node-1"},
				{Name: "node-2"},
				{Name: "node-3"},
			},
			ApiServerAddresses: []string{"https://api.gamma.example.com"},
		},
	}

	fmt.Printf("Inserting %d cluster documents...\n", len(clusters))
	if _, err := collection.InsertMany(ctx, clusters); err != nil {
		log.Fatalf("Failed to insert documents: %v", err)
	}

	fmt.Println("Successfully populated MongoDB with placeholder data.")
}
