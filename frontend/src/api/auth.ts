import { apiFetcher } from "../lib/apiClient";

export async function login(email: string, password: string): Promise<any> {
  try {
    const response = await apiFetcher("/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });
    return response;
  } catch (error) {
    throw error;
  }
}
