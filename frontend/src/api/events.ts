import { apiFetcher } from "@/lib/apiClient";
import type {
  CreateEventRequest,
  Event,
  UpdateEventRequest,
} from "@/types/event";

export async function getEvents(): Promise<Event[]> {
  const response = await apiFetcher("/api/events", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response.events ?? [];
}

export async function getEventById(id: string): Promise<Event> {
  const response = await apiFetcher(`/api/events/${id}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return response;
}

export async function createEvent(
  eventRequest: CreateEventRequest,
): Promise<any> {
  const response = await apiFetcher("/api/events", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(eventRequest),
  });
  return response;
}

export async function updateEvent(
  id: string,
  eventRequest: UpdateEventRequest,
): Promise<any> {
  const response = await apiFetcher(`/api/events/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(eventRequest),
  });
  return response;
}

export async function deleteEvent(id: string): Promise<any> {
  const response = await apiFetcher(`/api/events/${id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response;
}
