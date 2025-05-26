package controller

import (
	"net/http"
	"suleman37/Golang_Training/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	service.HandleLoginResponse(c, credentials.Email, credentials.Password)
}
