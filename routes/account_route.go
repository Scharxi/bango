package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/controller"
)

func AddAccountRoutes(router *gin.RouterGroup) {
	router.GET("/account", controller.GetAccount)
	router.POST("/account", controller.CreateAccount)
	router.PUT("/account/:id", controller.UpdateAccount)
	router.DELETE("/account/:id", controller.DeleteAccount)
}
