package controller

import (
	"net/http"
	// "suleman37/Golang_Training/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var credentials struct {
		Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
}