package main

import (
	"go-mongo/initializers"
	"go-mongo/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	if os.Getenv("ENV") != "production" {
		initializers.LoadEnv()
	}
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
		"message": "Hello from the other side!",
		})
	})

	routes.UserRoute(router)

	router.Run()
}