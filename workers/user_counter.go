package workers

import (
	"context"
	"log"
	"time"

	"golang-test/config"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func UserCount() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				countUsers()
			}
		}
	}()
}

func countUsers() {
	collection := config.GetDB().Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Error counting users: %v", err)
		return
	}

	log.Printf("Current number of users in database: %d", count)
}