package service

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FetchAllMessages() ([]map[string]interface{}, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Database connection error:", err)
		return nil, err
	}

	collection := client.Database("Golang").Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching users:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []map[string]interface{}
	if err = cursor.All(ctx, &users); err != nil {
		log.Println("Error decoding users:", err)
		return nil, err
	}
	for _, user := range users {
		delete(user, "password")
	}

	return users, nil
}