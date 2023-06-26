package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/controller"
)

func AddSampleRoutes(group *gin.RouterGroup) {
	group.GET("/hello", controller.SayHello)
}
