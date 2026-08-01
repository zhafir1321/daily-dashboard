import type { RegisterReq } from "@/components/RegisterView";
import { apiFetcher } from "../lib/apiClient";

export async function login(email: string, password: string): Promise<any> {
  const response = await apiFetcher("/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });
  return response;
}

export async function register(registerReq: RegisterReq): Promise<any> {
  const response = await apiFetcher("/auth/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(registerReq),
  });
  return response;
}
