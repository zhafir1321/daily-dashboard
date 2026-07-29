type ApiFetcherOptions = {
  method?: string;
  headers?: Record<string, string>;
  body?: any;
};

export async function apiFetcher(
  url: string,
  options: ApiFetcherOptions,
): Promise<any> {
  const token = localStorage.getItem("token");
  if (token) {
    options.headers = {
      ...options.headers,
      Authorization: `Bearer ${token}`,
    };
  }
  const response = await fetch(
    `${import.meta.env.VITE_API_URL}${url}`,
    options,
  );

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.message || "API request failed");
  }

  return data;
}
