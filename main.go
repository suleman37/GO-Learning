package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Model string `form:"model" json:"model"`
}

func main() {
	r := gin.Default()
	r.POST("/postdata", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"model": user.Model,
		})
	})

	r.Run(":8080")
}