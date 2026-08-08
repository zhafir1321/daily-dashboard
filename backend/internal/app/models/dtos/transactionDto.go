package dtos

import "backend/internal/app/entities"

type TransactionResponse struct {
	ID              int    `json:"id"`
	Type            string `json:"type"`
	Amount          string `json:"amount"`
	Description     string `json:"description"`
	Category        string `json:"category"`
	TransactionDate string `json:"transaction_date"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type GetTransactionResponse struct {
	Transaction *TransactionResponse `json:"transaction"`
	Message     string               `json:"message"`
}

type GetAllTransactionsResponse struct {
	Transactions []*TransactionResponse `json:"transactions"`
	Message      string                 `json:"message"`
}

type TransactionSummaryData struct {
	TotalIncome  string `json:"total_income"`
	TotalExpense string `json:"total_expense"`
	Balance      string `json:"balance"`
}

type GetTransactionSummaryResponse struct {
	Message string                  `json:"message"`
	Data    *TransactionSummaryData `json:"data"`
}

type CreateTransactionRequest struct {
	Type            string `json:"type" binding:"required"`
	Amount          string `json:"amount" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Category        string `json:"category" binding:"required"`
	TransactionDate string `json:"transaction_date" binding:"required"`
}

type UpdateTransactionRequest struct {
	Type            string `json:"type" binding:"omitempty"`
	Amount          string `json:"amount" binding:"omitempty"`
	Description     string `json:"description" binding:"omitempty"`
	Category        string `json:"category" binding:"omitempty"`
	TransactionDate string `json:"transaction_date" binding:"omitempty"`
}

type CreateTransactionResponse struct {
	Message string `json:"message"`
}

type UpdateTransactionResponse struct {
	Message string `json:"message"`
}

type DeleteTransactionResponse struct {
	Message string `json:"message"`
}

func (r *GetAllTransactionsResponse) MapTransactionsResponse(transactions []*entities.Transaction) {
	r.Transactions = []*TransactionResponse{}
	for _, transaction := range transactions {
		transactionResponse := &TransactionResponse{
			ID:              transaction.ID,
			Type:            transaction.Type,
			Amount:          transaction.Amount,
			Description:     transaction.Description,
			Category:        transaction.Category,
			TransactionDate: transaction.TransactionDate,
			CreatedAt:       transaction.CreatedAt,
			UpdatedAt:       transaction.UpdatedAt,
		}
		r.Transactions = append(r.Transactions, transactionResponse)
	}
}

func (r *TransactionResponse) MapTransactionResponse(transaction *entities.Transaction) {
	r.ID = transaction.ID
	r.Type = transaction.Type
	r.Amount = transaction.Amount
	r.Description = transaction.Description
	r.Category = transaction.Category
	r.TransactionDate = transaction.TransactionDate
	r.CreatedAt = transaction.CreatedAt
	r.UpdatedAt = transaction.UpdatedAt
}

func (r *CreateTransactionRequest) ToTransaction(userID int) *entities.Transaction {
	return &entities.Transaction{
		UserID:          userID,
		Type:            r.Type,
		Amount:          r.Amount,
		Description:     r.Description,
		Category:        r.Category,
		TransactionDate: r.TransactionDate,
	}
}
