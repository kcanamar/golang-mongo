package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"go-mongo/initializers"
)

func init() {
	if os.Getenv("ENV") != "production" {
		initializers.LoadEnv()
	}
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello from the other side!",
		})
	})
	router.Run()
}