package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
}

func main() {
	r := gin.Default()

	books := []Book{
		{ID: 1, Name: "Toyota Raize", Model: "2024"},
		{ID: 2, Name: "Grandee X", Model: "2025"},
	}

	r.GET("/cars", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Data Posted Sucessfully!",
		})
	})
	r.Run(":8080")
}
