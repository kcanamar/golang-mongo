package routes

import (
	"github.com/gin-gonic/gin"
	"go-mongo/controllers"
)

func UserRoute(router *gin.Engine){
	// Create
	router.POST("/users", controllers.CreateUser())
	// Index
	router.GET("/users", controllers.GetAllUsers())
	// Show
	router.GET("/users/:userId", controllers.GetUser())
}