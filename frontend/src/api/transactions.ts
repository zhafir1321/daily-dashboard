import { apiFetcher } from "@/lib/apiClient";
import type {
  CreateTransactionRequest,
  SummaryResponse,
  Transaction,
  UpdateTransactionRequest,
} from "@/types/transaction";

export async function getTransactions(): Promise<Transaction[]> {
  const response = await apiFetcher("/api/transactions", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return response.transactions ?? [];
}

export async function getSummary(): Promise<SummaryResponse> {
  const response = await apiFetcher("/api/transactions/summary", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return (
    response.data ?? {
      total_income: "0.00",
      total_expense: "0.00",
      balance: "0.00",
    }
  );
}

export async function getTransactionById(id: string): Promise<Transaction> {
  const response = await apiFetcher(`/api/transactions/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return response.transaction;
}

export async function createTransaction(
  transactionRequest: CreateTransactionRequest,
): Promise<any> {
  const response = await apiFetcher("/api/transactions", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(transactionRequest),
  });

  return response;
}

export async function updateTransaction(
  transactionId: string,
  updateRequest: UpdateTransactionRequest,
): Promise<any> {
  const response = await apiFetcher(`/api/transactions/${transactionId}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(updateRequest),
  });
  return response;
}

export async function deleteTransaction(transactionId: string): Promise<any> {
  const response = await apiFetcher(`/api/transactions/${transactionId}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}
