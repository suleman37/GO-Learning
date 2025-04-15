package service

import (
	"net/http"
	"suleman37/Golang_Training/middleware"
	"github.com/gin-gonic/gin"
)

func HandleLoginResponse(c *gin.Context, inputUsername, inputPassword, correctUsername, correctPassword string) {
	if inputUsername == correctUsername && inputPassword == correctPassword {
		token, err := middleware.CreateToken(correctUsername)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error geneating token "})
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}