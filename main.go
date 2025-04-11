package main

import (
	"net/http"
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Car struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`
}

func main() {

	file, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	filecontent := string(file)

	fmt.Println(filecontent)

	r := gin.Default()


	r.GET("/cars", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"fileContent": filecontent,
		})
	})

	r.POST("/create-cars", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Data Posted Sucessfully!",
		})
	})

	r.DELETE("/delete-cars", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Data Deleted Sucessfully!",
		})
	})

	r.PUT("/update-cars", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Data Updated Sucessfully!",
		})
	})

	r.Run(":8080")
}