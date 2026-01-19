package main

import (
	"log"
	"os"

	"github.com/dana-team/axiom/backend/internal/routes"
	"github.com/dana-team/axiom/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// @title           Axiom Backend API
// @version         1.0
// @description     Backend service for Axiom cluster management.
// @BasePath        /v1

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load .env file")
	}

	logger := initializeLogger()
	defer syncLogger(logger)

	mongoClient, err := utils.InitMongoDB()
	if err != nil {
		logger.Fatal("Failed to initialize MongoDB client", zap.Error(err))
	}
	defer func(mongoClient *utils.MongoClient) {
		err := mongoClient.Disconnect()
		if err != nil {
			logger.Fatal("Failed to disconnect from MongoDB", zap.Error(err))
		}
	}(mongoClient)

	router := gin.Default()
	routes.SetupClusterRoutes(router, mongoClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Starting server", zap.String("port", port))
	if err := router.Run(":" + port); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func initializeLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logger
}

func syncLogger(logger *zap.Logger) {
	if err := logger.Sync(); err != nil {
		log.Fatalf("Failed to sync logger: %v", err)
	}
}
