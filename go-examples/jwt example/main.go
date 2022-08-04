package main

import (
	"jwt-example/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/welcome", handler.Welcome)
	route.POST("/signin", handler.Signin)
	route.Run("localhost:8000")
}
