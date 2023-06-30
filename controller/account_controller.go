package controller

import (
	"github.com/gin-gonic/gin"
	db "go-gin-template/db/sqlc"
	"go-gin-template/util"
	"net/http"
)

func CreateAccount(c *gin.Context) {
	var args db.CreateBankAccountParams
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountNumber, err := util.GenerateAccountNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := util.GetStore().CreateBankAccountTx(c, args, accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Account created successfully"})
}

func OpenAccount(c *gin.Context) {

}

func GetAccount(c *gin.Context) {

}

func UpdateAccount(c *gin.Context) {

}

func DeleteAccount(c *gin.Context) {

}
