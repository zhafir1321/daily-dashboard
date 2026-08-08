package entities

type Transaction struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Type            string `json:"type"`
	Amount          string `json:"amount"`
	Description     string `json:"description"`
	Category        string `json:"category"`
	TransactionDate string `json:"transaction_date"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
