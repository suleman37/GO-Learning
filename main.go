package main

import (
	"fmt"
	"suleman37/Golang_Training/controller"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default();
	r.POST("/register", controller.Login)
	r.POST("/login", controller.Login)
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

func verifyToken(tokenString string, c *gin.Context) error {
	secretKey := []byte("123456789")
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