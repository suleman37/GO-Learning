package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func HandleRegisterResponse(c *gin.Context, email, username, password string) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	log.Println("Connected to MongoDB successfully")

	if isEmailRegistered(client, email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	err = saveUserToDatabase(client, email, username, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user to database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func isEmailRegistered(client *mongo.Client, email string) bool {
	collection := client.Database("Golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result bson.M
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	return err == nil
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
};