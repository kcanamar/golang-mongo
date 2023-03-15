package routes

import (
	"github.com/gin-gonic/gin"
	"go-mongo/controllers"
)

func UserRoute(router *gin.Engine){
	router.POST("/user", controllers.CreateUser())
}