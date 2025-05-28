package main

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"suleman37/Golang_Training/controller"
	"suleman37/Golang_Training/dbconnect"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.Use(corsMiddleware())
	r.POST("/register", controller.Register)
	r.GET("/users", controller.GetAllUsers)
	r.POST("/login", controller.Login)
	r.POST("/message", controller.CreateMessage)
	dbconnect.DBConnection()

	r.POST("/token", func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(400, gin.H{"message": "Bearer token is missing"})
			return
		}
		if err := verifyToken(tokenString, c); err != nil {
			return
		}
		c.JSON(200, gin.H{"message": "Valid Token"})
	})

	r.Run(":8000")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func verifyToken(tokenString string, c *gin.Context) error {
	secretKey := []byte("123abcxyz")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		return err
	}
	if !token.Valid {
		c.JSON(401, gin.H{"error": "Invalid token"})
	}
	return nil
}
