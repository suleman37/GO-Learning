package controller

import (
	"net/http"
	"suleman37/Golang_Training/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	service.HandleRegisterResponse(c, newUser.Email, newUser.Username, newUser.Password)
}