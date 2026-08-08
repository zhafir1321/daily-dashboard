package services

import (
	"backend/internal/app/models"
	"backend/internal/app/models/dtos"
	"backend/internal/app/repositories"
	"database/sql"
	"errors"
	"net/http"

	"github.com/rs/zerolog"
)

type Transaction struct {
	transactionRepository *repositories.Transaction
	logger                zerolog.Logger
}

func NewTransactionService(transactionRepository *repositories.Transaction, logger zerolog.Logger) *Transaction {
	return &Transaction{transactionRepository: transactionRepository, logger: logger}
}

func (ts *Transaction) GetSummary(userID int) (*dtos.GetTransactionSummaryResponse, *models.ErrorResponse) {
	response := &dtos.GetTransactionSummaryResponse{}
	response.Data = &dtos.TransactionSummaryData{}

	totalIncome, err := ts.transactionRepository.GetSumOfTransactionsByType(userID, "income")
	if err != nil {
		ts.logger.Error().Err(err).Msg("Failed to retrieve total income")
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve total income",
		}
	}
	response.Data.TotalIncome = totalIncome

	totalExpense, err := ts.transactionRepository.GetSumOfTransactionsByType(userID, "expense")
	if err != nil {
		ts.logger.Error().Err(err).Msg("Failed to retrieve total expense")
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve total expense",
		}
	}
	response.Data.TotalExpense = totalExpense

	balance, err := ts.transactionRepository.GetBalance(userID)
	if err != nil {
		ts.logger.Error().Err(err).Msg("Failed to retrieve balance")
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve balance",
		}
	}
	response.Data.Balance = balance

	response.Message = "Transaction summary retrieved successfully"

	return response, nil
}

func (ts *Transaction) GetAllTransactions(userID int) (*dtos.GetAllTransactionsResponse, *models.ErrorResponse) {
	response := &dtos.GetAllTransactionsResponse{}
	response.Transactions = []*dtos.TransactionResponse{}

	queriedTransactions, err := ts.transactionRepository.GetAllTransactions(userID)
	if err != nil {
		ts.logger.Error().Err(err).Msg("Failed to retrieve transactions")
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve transactions",
		}
	}

	response.MapTransactionsResponse(queriedTransactions)
	response.Message = "Transactions retrieved successfully"
	return response, nil
}

func (ts *Transaction) GetTransaction(transactionID int, userID int) (*dtos.GetTransactionResponse, *models.ErrorResponse) {
	response := &dtos.GetTransactionResponse{}
	response.Transaction = &dtos.TransactionResponse{}
	queriedTransaction, queryErr := ts.transactionRepository.FindByID(transactionID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Transaction not found",
			}
		}

		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve transaction",
		}
	}

	response.Transaction.MapTransactionResponse(queriedTransaction)
	response.Message = "Transaction retrieved successfully"
	return response, nil
}

func (ts *Transaction) CreateTransaction(transactionReq *dtos.CreateTransactionRequest, userID int) (*dtos.CreateTransactionResponse, *models.ErrorResponse) {
	transaction := transactionReq.ToTransaction(userID)

	if err := ts.transactionRepository.CreateTransaction(transaction); err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create transaction",
		}
	}

	response := &dtos.CreateTransactionResponse{
		Message: "Transaction created successfully",
	}

	return response, nil
}

func (ts *Transaction) UpdateTransaction(transactionID int, transactionReq *dtos.UpdateTransactionRequest, userID int) (*dtos.UpdateTransactionResponse, *models.ErrorResponse) {
	existingTransaction, queryErr := ts.transactionRepository.FindByID(transactionID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Transaction not found",
			}
		}

		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve transaction",
		}
	}

	if transactionReq.Type != "" {
		existingTransaction.Type = transactionReq.Type
	}

	if transactionReq.Amount != "" {
		existingTransaction.Amount = transactionReq.Amount
	}

	if transactionReq.Description != "" {
		existingTransaction.Description = transactionReq.Description
	}

	if transactionReq.Category != "" {
		existingTransaction.Category = transactionReq.Category
	}

	if transactionReq.TransactionDate != "" {
		existingTransaction.TransactionDate = transactionReq.TransactionDate
	}

	err := ts.transactionRepository.UpdateTransaction(existingTransaction)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update transaction",
		}
	}

	response := &dtos.UpdateTransactionResponse{
		Message: "Transaction updated successfully",
	}

	return response, nil
}

func (ts *Transaction) DeleteTransaction(transactionID int, userID int) (*dtos.DeleteTransactionResponse, *models.ErrorResponse) {
	existingTransaction, queryErr := ts.transactionRepository.FindByID(transactionID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Transaction not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve transaction",
		}
	}

	err := ts.transactionRepository.DeleteTransaction(existingTransaction.ID, userID)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete transaction",
		}
	}

	response := &dtos.DeleteTransactionResponse{
		Message: "Transaction deleted successfully",
	}

	return response, nil
}
