package main

import (
	"cookie-example/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)
	r.Run("localhost:8000")

}
