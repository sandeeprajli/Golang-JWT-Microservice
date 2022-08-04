package main

import (
	"google-auth-example/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", handlers.Login)
	r.GET("/callback", handlers.Callback)
	r.Run("127.0.0.1:8000")
}
