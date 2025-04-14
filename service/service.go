package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLoginResponse(c *gin.Context, inputUsername, inputPassword, correctUsername, correctPassword string) {
	if inputUsername == correctUsername && inputPassword == correctPassword {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}