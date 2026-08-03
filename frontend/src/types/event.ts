export type Event = {
  id: string;
  title: string;
  description: string;
  event_date: string;
  event_time: string | null;
  created_at: string;
  updated_at: string;
};

export type CreateEventRequest = {
  title: string;
  description: string;
  event_date: string;
  event_time?: string | null;
};

export type UpdateEventRequest = {
  title?: string;
  description?: string;
  event_date?: string;
  event_time?: string | null;
};
