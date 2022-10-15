package controllers

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"ottoTest/models"
	"ottoTest/utils/token"
	"io/ioutil"
	"encoding/json"
    "fmt"
    "log"
)

// returns current user balance
func GetUserBalance(c *gin.Context){

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	balance,err := models.ReturnBalance(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","Balance":balance})
}

type TopUpInput struct {
	Debit float64 `json:"debit" binding:"required"`
}

// top up balance
func TopUp(c *gin.Context){

	var input TopUpInput

	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Token error": err.Error()})
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := models.Transaction{}

	t.Debit = input.Debit
	t.WalletID = user_id

	_,err = t.CreateTransaction()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_,err = models.UpdateBalance(user_id, input.Debit, 0)

	c.JSON(http.StatusOK, gin.H{"message":"Top Up success"})

}

// returns all user transactions
func GetUserTransactions(c *gin.Context){

	user_id, err := token.ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	transaction,err := models.GetTransactions(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"success","Transactions":transaction})
}



type BillerTransaction struct {
    Id int `json:"id"`
	Category string `json:"category"` 
	Product string `json:"product"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Fee float64 `json:"fee"`
}

type BillerTransactionResponse struct {
    Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data BillerTransaction `json:"data"`
}

type BillerInput struct {
	ID string `json:"id" binding:"required"`
}

type BillerTransactionsResponse struct {
    Code int `json:"code"`
	Status string `json:"status"`
	Message string `json:"message"`
	Data []BillerTransaction `json:"data"`
}
// returns single biller transaction
func ConfirmBillerTransaction(c *gin.Context) {

	var input BillerInput


	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := fmt.Sprintf("https://phoenix-imkas.ottodigital.id/interview/biller/v1/detail?billerId=%s", input.ID)

	response, err := http.Get(url)
    if err != nil {
        fmt.Print(err.Error())
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject BillerTransactionResponse
    json.Unmarshal(responseData, &responseObject)

	c.JSON(http.StatusOK, gin.H{"message":"success","Transactions":responseObject.Data})
}

// returns all biller transactions
func GetBillerTransactions(c *gin.Context){

    response, err := http.Get("https://phoenix-imkas.ottodigital.id/interview/biller/v1/list")
    if err != nil {
        fmt.Print(err.Error())
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    var responseObject BillerTransactionsResponse
    json.Unmarshal(responseData, &responseObject)

	c.JSON(http.StatusOK, gin.H{"message":"success","Transactions":responseObject.Data})}