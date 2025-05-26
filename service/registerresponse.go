package service

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleRegisterResponse(c *gin.Context, email, username, password string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	log.Println("Connected to MongoDB successfully")

	err = saveUserToDatabase(client, email, username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user to database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func saveUserToDatabase(client *mongo.Client, email, username, password string) error {
	collection := client.Database("Golang").Collection("users")

	user := map[string]string{
		"email":    email,
		"username": username,
		"password": password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Error inserting user into MongoDB:", err)
		return err
	}

	return nil
}