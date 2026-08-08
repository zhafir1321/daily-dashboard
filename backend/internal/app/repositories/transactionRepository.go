package repositories

import (
	"backend/internal/app/database"
	"backend/internal/app/entities"
	"database/sql"
)

type Transaction struct {
	database.BaseSQLRepository[entities.Transaction]
}

func NewTransactionRepository(db *sql.DB) *Transaction {
	return &Transaction{
		BaseSQLRepository: database.BaseSQLRepository[entities.Transaction]{DB: db},
	}
}

func mapTransaction(rows *sql.Row, t *entities.Transaction) error {
	return rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Amount, &t.Description, &t.Category, &t.TransactionDate, &t.CreatedAt, &t.UpdatedAt)
}

func mapTransactions(rows *sql.Rows, t *entities.Transaction) error {
	return rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Amount, &t.Description, &t.Category, &t.TransactionDate, &t.CreatedAt, &t.UpdatedAt)
}

func (r *Transaction) FindByID(id int, userID int) (*entities.Transaction, error) {
	return r.SelectSingle(
		mapTransaction,
		"SELECT id, user_id, type, amount, description, category, TO_CHAR(transaction_date, 'YYYY-MM-DD') as transaction_date, created_at, updated_at FROM transactions WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)
}

func (r *Transaction) GetAllTransactions(userID int) ([]*entities.Transaction, error) {
	return r.SelectMultiple(
		mapTransactions,
		"SELECT id, user_id, type, amount, description, category, TO_CHAR(transaction_date, 'YYYY-MM-DD') as transaction_date, created_at, updated_at FROM transactions WHERE user_id = $1",
		userID,
	)
}

func (r *Transaction) GetTransactionsByType(userID int, transactionType string) ([]*entities.Transaction, error) {
	return r.SelectMultiple(
		mapTransactions,
		"SELECT id, user_id, type, amount, description, category, TO_CHAR(transaction_date, 'YYYY-MM-DD') as transaction_date, created_at, updated_at FROM transactions WHERE user_id = $1 AND type = $2",
		userID,
		transactionType,
	)
}

func (r *Transaction) GetSumOfTransactionsByType(userID int, transactionType string) (string, error) {
	var sum string
	err := r.DB.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = $1 AND type = $2", userID, transactionType).Scan(&sum)
	return sum, err
}

func (r *Transaction) GetSumOfTransactionsByCategory(userID int, category string) (string, error) {
	var sum string
	err := r.DB.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = $1 AND category = $2", userID, category).Scan(&sum)
	return sum, err
}

func (r *Transaction) GetSumOfTransactionsByTypeAndCategory(userID int, transactionType string, category string) (string, error) {
	var sum string
	err := r.DB.QueryRow("SELECT COALESCE(SUM(amount), 0) FROM transactions WHERE user_id = $1 AND type = $2 AND category = $3", userID, transactionType, category).Scan(&sum)
	return sum, err
}

func (r *Transaction) GetBalance(userID int) (string, error) {
	var sum string
	err := r.DB.QueryRow("SELECT COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE -amount END), 0) FROM transactions WHERE user_id = $1", userID).Scan(&sum)
	return sum, err
}

func (r *Transaction) CreateTransaction(transaction *entities.Transaction) error {
	id, err := r.Insert(
		"INSERT INTO transactions (user_id, type, amount, description, category, transaction_date) VALUES ($1, $2, $3, $4, $5, $6)",
		transaction.UserID,
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Category,
		transaction.TransactionDate,
	)
	transaction.ID = id
	return err
}

func (r *Transaction) UpdateTransaction(transaction *entities.Transaction) error {
	_, err := r.ExecuteQuery(
		"UPDATE transactions SET type = $1, amount = $2, description = $3, category = $4, transaction_date = $5, updated_at = NOW() WHERE id = $6 AND user_id = $7",
		transaction.Type,
		transaction.Amount,
		transaction.Description,
		transaction.Category,
		transaction.TransactionDate,
		transaction.ID,
		transaction.UserID,
	)
	return err
}

func (r *Transaction) DeleteTransaction(id int, userID int) error {
	_, err := r.ExecuteQuery(
		"DELETE FROM transactions WHERE id = $1 AND user_id = $2",
		id,
		userID,
	)
	return err
}
