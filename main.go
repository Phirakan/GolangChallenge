package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"golang-test/config"
	"golang-test/routes"
	"golang-test/workers"
)

func main() {
	// Connect to MongoDB
	client, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect from MongoDB:", err)
		}
	}()

	// Create indexes
	if err := createIndexes(); err != nil {
		log.Fatal("Failed to create indexes:", err)
	}

	// Start background worker
	workers.UserCount()

	// Setup Gin router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

//Create index function
func createIndexes() error {
	collection := config.GetDB().Collection("user")
	ctx := context.Background()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	log.Println("Indexes created successfully")
	return nil
}