export type Transaction = {
  id: string;
  type: TransactionType;
  amount: string;
  description: string;
  category: string;
  transaction_date: string;
  created_at: string;
  updated_at: string;
};

export type TransactionType = "income" | "expense";

export type CreateTransactionRequest = {
  type: TransactionType;
  amount: string;
  description: string;
  category: string;
  transaction_date: string;
};

export type UpdateTransactionRequest = {
  type?: TransactionType;
  amount?: string;
  description?: string;
  category?: string;
  transaction_date?: string;
};

export type SummaryResponse = {
  total_income: string;
  total_expense: string;
  balance: string;
};
