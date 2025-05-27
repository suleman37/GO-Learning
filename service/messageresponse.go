package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func HandleMessageResponse(c *gin.Context, sender, receiver, chatRoomId, content string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	log.Println("Connected to MongoDB successfully")

	err = saveMessageToDatabase(client, sender, receiver, chatRoomId, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving message to database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func saveMessageToDatabase(client *mongo.Client, sender, receiver, chatRoomId, content string) error {
	collection := client.Database("Golang").Collection("messages")

	message := map[string]string{
		"sender":     sender,
		"receiver":   receiver,
		"chatRoomId": chatRoomId,
		"content":    content,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, message)
	if err != nil {
		log.Println("Error inserting message into MongoDB:", err)
		return err
	}

	return nil
}
