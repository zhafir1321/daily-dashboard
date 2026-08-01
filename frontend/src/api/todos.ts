import type { Todo } from "@/components/TodosView";
import { apiFetcher } from "../lib/apiClient";

export async function getTodos(): Promise<Todo[]> {
  const response = await apiFetcher("/api/todos", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  // console.log("Fetched todos:", response.todos); // Log the fetched todos for debugging
  return response.todos ?? [];
}

export async function getTodoById(id: string): Promise<Todo> {
  const response = await apiFetcher(`/api/todos/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}

export async function createTodo(
  title: string,
  description: string,
): Promise<any> {
  const response = await apiFetcher("/api/todos", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title, description }),
  });
  return response;
}

export async function updateTodo(
  id: string,
  title?: string,
  description?: string,
): Promise<any> {
  const response = await apiFetcher(`/api/todos/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title, description }),
  });
  return response;
}

export async function toggleTodoCompletion(id: string): Promise<any> {
  const response = await apiFetcher(`/api/todos/${id}/complete`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}

export async function toggleTodoIncompletion(id: string): Promise<any> {
  const response = await apiFetcher(`/api/todos/${id}/incomplete`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}

export async function deleteTodo(id: string): Promise<any> {
  const response = await apiFetcher(`/api/todos/${id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}
