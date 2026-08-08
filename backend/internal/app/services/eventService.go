package services

import (
	"backend/internal/app/models"
	"backend/internal/app/models/dtos"
	"backend/internal/app/repositories"
	"database/sql"
	"errors"
	"net/http"

	"github.com/rs/zerolog"
)

type Event struct {
	eventRepository *repositories.Event
	logger          zerolog.Logger
}

func NewEventService(eventRepository *repositories.Event, logger zerolog.Logger) *Event {
	return &Event{eventRepository: eventRepository, logger: logger}
}

func (es *Event) GetAllEvents(userID int) (*dtos.GetAllEventsResponse, *models.ErrorResponse) {
	response := &dtos.GetAllEventsResponse{}

	queriedEvents, err := es.eventRepository.GetAllEvents(userID)
	if err != nil {
		es.logger.Error().Err(err).Msg("Failed to retrieve events")
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve events",
		}
	}

	response.MapEventsResponse(queriedEvents)
	response.Message = "Events retrieved successfully"
	return response, nil
}

func (es *Event) GetEvent(eventID int, userID int) (*dtos.GetEventResponse, *models.ErrorResponse) {
	response := &dtos.GetEventResponse{}
	response.Event = &dtos.EventResponse{}
	queriedEvent, queryErr := es.eventRepository.FindByID(eventID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Event not found",
			}
		}

		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve event",
		}
	}

	response.Event.MapEventResponse(queriedEvent)
	response.Message = "Event retrieved successfully"
	return response, nil
}

func (es *Event) CreateEvent(eventReq *dtos.CreateEventRequest, userID int) (*dtos.CreateEventResponse, *models.ErrorResponse) {
	event := eventReq.ToEvent(userID)

	if err := es.eventRepository.CreateEvent(event); err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create event",
		}
	}

	response := &dtos.CreateEventResponse{
		Message: "Event created successfully",
	}

	return response, nil
}

func (es *Event) UpdateEvent(eventID int, eventReq *dtos.UpdateEventRequest, userID int) (*dtos.UpdateEventResponse, *models.ErrorResponse) {
	existingEvent, queryErr := es.eventRepository.FindByID(eventID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Event not found",
			}
		}

		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve event",
		}
	}

	if eventReq.Title != "" {
		existingEvent.Title = eventReq.Title
	}

	if eventReq.Description != "" {
		existingEvent.Description = eventReq.Description
	}

	if eventReq.EventDate != "" {
		existingEvent.EventDate = eventReq.EventDate
	}

	existingEvent.EventTime = eventReq.EventTime

	err := es.eventRepository.UpdateEvent(existingEvent)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update event",
		}
	}

	response := &dtos.UpdateEventResponse{
		Message: "Event updated successfully",
	}
	return response, nil
}

func (es *Event) DeleteEvent(eventID int, userID int) (*dtos.DeleteEventResponse, *models.ErrorResponse) {
	existingEvent, queryErr := es.eventRepository.FindByID(eventID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Event not found",
			}
		}

		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve event",
		}
	}

	err := es.eventRepository.DeleteEvent(existingEvent.ID, userID)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete event",
		}
	}

	response := &dtos.DeleteEventResponse{
		Message: "Event deleted successfully",
	}
	return response, nil
}
