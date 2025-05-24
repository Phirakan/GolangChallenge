package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database

func ConnectDB() (*mongo.Client, error) {

	mongoURL := os.Getenv("MONGODB_URL")
	if mongoURL == "" {
	
		mongoURL = "mongodb+srv://mosuuuutech:mosuuuutech@cluster0.pjni7b1.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURL).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	// Set database
	DB = client.Database("golang_challenge")

	return client, nil
}

func GetDB() *mongo.Database {
	return DB
}