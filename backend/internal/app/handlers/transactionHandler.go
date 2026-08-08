package handlers

import (
	"backend/internal/app/helpers"
	"backend/internal/app/models/dtos"
	"backend/internal/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Transaction struct {
	transactionService *services.Transaction
	logger             zerolog.Logger
}

func NewTransactionHandler(transactionService *services.Transaction, logger zerolog.Logger) *Transaction {
	return &Transaction{transactionService: transactionService, logger: logger}
}

func (h *Transaction) GetSummary(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)
	summary, err := h.transactionService.GetSummary(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}
	ctx.JSON(200, summary)
}

func (h *Transaction) GetTransactions(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)
	transactions, err := h.transactionService.GetAllTransactions(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}
	ctx.JSON(200, transactions)
}

func (h *Transaction) GetTransaction(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Transaction ID is not valid"})
		return
	}

	userID := helpers.GetUserId(ctx)
	transaction, transactionErr := h.transactionService.GetTransaction(transactionID, userID)
	if transactionErr != nil {
		ctx.AbortWithStatusJSON(transactionErr.Code, transactionErr)
		return
	}
	ctx.JSON(200, transaction)
}

func (h *Transaction) CreateTransaction(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)
	var createTransactionRequest dtos.CreateTransactionRequest

	if err := ctx.ShouldBindJSON(&createTransactionRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	newTransaction, err := h.transactionService.CreateTransaction(&createTransactionRequest, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}
	ctx.JSON(200, newTransaction)
}

func (h *Transaction) UpdateTransaction(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Transaction ID is not valid"})
		return
	}

	userID := helpers.GetUserId(ctx)
	var updateTransactionRequest dtos.UpdateTransactionRequest

	if err := ctx.ShouldBindJSON(&updateTransactionRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedTransaction, updateTransactionErr := h.transactionService.UpdateTransaction(transactionID, &updateTransactionRequest, userID)
	if updateTransactionErr != nil {
		ctx.AbortWithStatusJSON(updateTransactionErr.Code, updateTransactionErr)
		return
	}
	ctx.JSON(200, updatedTransaction)
}

func (h *Transaction) DeleteTransaction(ctx *gin.Context) {
	transactionID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Transaction ID is not valid"})
		return
	}

	userID := helpers.GetUserId(ctx)
	deleteTransactionResponse, deleteTransactionErr := h.transactionService.DeleteTransaction(transactionID, userID)
	if deleteTransactionErr != nil {
		ctx.AbortWithStatusJSON(deleteTransactionErr.Code, deleteTransactionErr)
		return
	}
	ctx.JSON(200, deleteTransactionResponse)
}
