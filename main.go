package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	ID     int    `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
}

func main() {
	r := gin.Default()

	books := []Book{
		{ID: 1, Name: "Toyota Raiz", Model: "2024"},
		{ID: 2, Name: "Grandee X", Model: "2025"},
	}

	r.GET("/cars", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})

	r.Run(":8080")
}