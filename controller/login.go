package controller

import (
	"net/http"
	"suleman37/Golang_Training/service"

	"github.com/gin-gonic/gin"
)

var username = "admin"
var password = "password123"

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	service.HandleLoginResponse(c, credentials.Username, credentials.Password, username, password)
}