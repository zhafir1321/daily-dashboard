package dtos

import "backend/internal/app/entities"

type EventResponse struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	EventDate   string  `json:"event_date"`
	EventTime   *string `json:"event_time"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetEventResponse struct {
	Event   *EventResponse `json:"event"`
	Message string         `json:"message"`
}

type GetAllEventsResponse struct {
	Events  []*EventResponse `json:"events"`
	Message string           `json:"message"`
}

type CreateEventRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	EventDate   string  `json:"event_date" binding:"required"`
	EventTime   *string `json:"event_time" binding:"omitempty"`
}

type UpdateEventRequest struct {
	Title       string  `json:"title" binding:"omitempty"`
	Description string  `json:"description" binding:"omitempty"`
	EventDate   string  `json:"event_date" binding:"omitempty"`
	EventTime   *string `json:"event_time" binding:"omitempty"`
}

type CreateEventResponse struct {
	Message string `json:"message"`
}

type DeleteEventResponse struct {
	Message string `json:"message"`
}

type UpdateEventResponse struct {
	Message string `json:"message"`
}

func (r *GetAllEventsResponse) MapEventsResponse(events []*entities.Event) {
	r.Events = []*EventResponse{}
	for _, event := range events {
		eventResponse := &EventResponse{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			EventDate:   event.EventDate,
			EventTime:   event.EventTime,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
		}
		r.Events = append(r.Events, eventResponse)
	}
}

func (r *EventResponse) MapEventResponse(event *entities.Event) {
	r.ID = event.ID
	r.Title = event.Title
	r.Description = event.Description
	r.EventDate = event.EventDate
	r.EventTime = event.EventTime
	r.CreatedAt = event.CreatedAt
	r.UpdatedAt = event.UpdatedAt
}

func (er *CreateEventRequest) ToEvent(userID int) *entities.Event {
	return &entities.Event{
		Title:       er.Title,
		Description: er.Description,
		EventDate:   er.EventDate,
		EventTime:   er.EventTime,
		UserID:      userID,
	}
}
