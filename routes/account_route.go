package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-template/controller"
)

func AddAccountRoutes(router *gin.RouterGroup) {
	router.GET("/account/:accnum", controller.GetAccount)
	router.GET("/account/:accnum/holder", controller.GetAccountWithHolder)
	router.GET("/account/holder/:id", controller.GetAccountsWithHolderId)
	router.GET("/account/:accnum/validate", controller.ValidateAccountNumber)
	router.POST("/account", controller.CreateAccount)
	router.POST("/account/:holderid/open", controller.OpenAccount)
	router.PUT("/account/:id", controller.UpdateAccount)
	router.DELETE("/account/:id", controller.DeleteAccount)
}
