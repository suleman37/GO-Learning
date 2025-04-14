package main

import (
	"suleman37/Golang_Training/controller"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/postdata", controller.Login)

	r.Run(":8000")
}