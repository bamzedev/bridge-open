package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Validator1PrivateKey         string
	Validator2PrivateKey         string
	BridgeContractAddressRopsten string
	BridgeContractAddressRinkeby string
	InfuraKey                    string
}

type Transaction struct {
	Signatures             [2]string `json:"signatures"`
	DepositTransactionHash string    `json:"depositTransactionHash"`
	IsBurn                 bool      `json:"isBurn"`
	ClaimTransactionHash   string    `json:"claimTransactionHash"`
	TokenAddress           string    `json:"tokenAddress"`
	Symbol                 string    `json:"symbol"`
	Name                   string    `json:"name"`
	Amount                 string    `json:"amount"`
	Recipient              string    `json:"recipient"`
	FromChainId            string    `json:"fromChainId"`
	ToChainId              string    `json:"toChainId"`
	Date                   int64     `json:"date"`
	Claimed                bool      `json:"claimed"`
}

var Transactions = make(map[string]Transaction)
var configuration = Configuration{}

func getTransactions(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Transactions)
}

func getTransactionByHash(tx string) (*Transaction, error) {
	if val, ok := Transactions[tx]; ok {
		return &val, nil
	}
	return nil, errors.New("no transaction with that hash")
}

func getTransactionsByAddress(context *gin.Context) {
	result := []Transaction{}
	for _, v := range Transactions {
		if strings.EqualFold(v.Recipient, context.Param("address")) {
			result = append(result, v)
		}
	}
	context.IndentedJSON(http.StatusOK, result)
}

func getTransaction(context *gin.Context) {
	tx := context.Param("transaction")
	transaction, err := getTransactionByHash(tx)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
	}
	context.IndentedJSON(http.StatusOK, transaction)
}

func updateTransaction(context *gin.Context) {
	tx, err := getTransactionByHash(context.Param("transaction"))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found"})
	}

	if err := context.BindJSON(&tx); err != nil {
		return
	}
	updatedTransaction, err := UpdateClaimed(*tx)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err)
	}
	Transactions[updatedTransaction.DepositTransactionHash] = updatedTransaction
	context.IndentedJSON(http.StatusOK, updatedTransaction)
}

func addTransaction(context *gin.Context) {
	var newTransaction Transaction

	if err := context.BindJSON(&newTransaction); err != nil {
		return
	}
	_, err := getTransactionByHash(newTransaction.DepositTransactionHash)
	if err == nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "transaction already added"})
		return
	}
	signedTransaction, err := Sign(newTransaction)
	if err != nil {
		panic(err)
	}
	Transactions[signedTransaction.DepositTransactionHash] = signedTransaction
	context.IndentedJSON(http.StatusCreated, signedTransaction)
}

func config() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
}

func Run() {

	config()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PATCH"},
		AllowHeaders: []string{"*"},
	}))
	router.GET("/transactions", getTransactions)
	router.POST("/transactions", addTransaction)
	router.GET("/transactions/:transaction", getTransaction)
	router.PATCH("/transactions/:transaction", updateTransaction)
	router.GET("/transactions/address/:address", getTransactionsByAddress)
	router.Run("localhost:8080")
}
