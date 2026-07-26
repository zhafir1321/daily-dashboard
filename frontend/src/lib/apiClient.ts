type ApiFetherOptions = {
  method?: string;
  headers?: Record<string, string>;
  body?: any;
};

export async function apiFetcher(
  url: string,
  options: ApiFetherOptions,
): Promise<any> {
  const response = await fetch(
    `${import.meta.env.VITE_API_URL}${url}`,
    options,
  );
  return response.json();
}
