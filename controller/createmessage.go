package controller

import (
	"net/http"
	"suleman37/Golang_Training/service"

	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	var message struct {
		Sender    string `json:"sender"`
		Receiver  string `json:"receiver"`
		ChatRoomId string `json:"chatRoomId"`
		Content   string `json:"content"`
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	service.HandleMessageResponse(c, message.Sender, message.Receiver, message.ChatRoomId, message.Content)
}