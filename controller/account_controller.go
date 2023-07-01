package controller

import (
	"github.com/gin-gonic/gin"
	db "go-gin-template/db/sqlc"
	"go-gin-template/util"
	"log"
	"net/http"
	"strconv"
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

func ValidateAccountNumber(c *gin.Context) {
	param := c.Param("accnum")
	accountNumber, err := strconv.ParseInt(param, 10, 64)

	exists, err := util.GetStore().DoesAccountNumberExist(c, accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

// OpenAccount POST /account/:holderid/open
// It opens a new account for a given account holder
func OpenAccount(c *gin.Context) {
	param := c.Param("holderid")
	holderId, err := strconv.ParseInt(param, 10, 64)

	log.Println("Holder ID: ", holderId)

	accountNumber, err := util.GenerateAccountNumber()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := util.GetStore().OpenBankAccountTx(c, int32(holderId), accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Successfully opened a new bank account"})
}

// GetAccount GET /account/:accnum
// Retrieve a single account by its account number
func GetAccount(c *gin.Context) {
	param := c.Param("accnum")
	accountNumber, err := strconv.ParseInt(param, 10, 64)

	account, err := util.GetStore().GetAccountByAccountNumber(c, accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// GetAccountWithHolder GET /account/:accnum/holder
// Retrieve a single account including the account holder information by its account number
func GetAccountWithHolder(c *gin.Context) {
	param := c.Param("accnum")
	accountNumber, err := strconv.ParseInt(param, 10, 64)

	account, err := util.GetStore().GetAccountWithHolder(c, accountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

// GetAccountsWithHolderId GET /account/holder/:id
// Retrieve all accounts including by its account holder id
func GetAccountsWithHolderId(c *gin.Context) {
	var limit, offset int32

	util.GetPaging(c, &offset, &limit)

	param := c.Param("id")
	holderId, err := strconv.ParseInt(param, 10, 64)

	var args = db.GetAccountsFromHolderParams{
		AccountHolderID: int32(holderId),
		Limit:           limit,
		Offset:          offset,
	}

	accounts, err := util.GetStore().GetAccountsFromHolder(c, args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

func UpdateAccount(c *gin.Context) {

}

func DeleteAccount(c *gin.Context) {

}
