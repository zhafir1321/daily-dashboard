package handlers

import (
	"backend/internal/app/helpers"
	"backend/internal/app/models/dtos"
	"backend/internal/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Event struct {
	eventService *services.Event
	logger       zerolog.Logger
}

func NewEventHandler(eventService *services.Event, logger zerolog.Logger) *Event {
	return &Event{eventService: eventService, logger: logger}
}

func (h *Event) GetAllEvents(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)

	allEvents, err := h.eventService.GetAllEvents(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}

	ctx.JSON(200, allEvents)
}

func (h *Event) GetEvent(ctx *gin.Context) {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Event ID is not valid"})
		return
	}

	userID := helpers.GetUserId(ctx)
	event, eventErr := h.eventService.GetEvent(eventID, userID)
	if eventErr != nil {
		ctx.AbortWithStatusJSON(eventErr.Code, eventErr)
		return
	}

	ctx.JSON(200, event)
}

func (h *Event) CreateEvent(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)
	var createEventRequest dtos.CreateEventRequest

	if err := ctx.ShouldBindJSON(&createEventRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	newEvent, err := h.eventService.CreateEvent(&createEventRequest, userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}

	ctx.JSON(201, newEvent)
}

func (h *Event) UpdateEvent(ctx *gin.Context) {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Event ID is not valid"})
		return
	}

	userID := helpers.GetUserId(ctx)

	var updateEventRequest dtos.UpdateEventRequest
	if err := ctx.ShouldBindJSON(&updateEventRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	event, eventErr := h.eventService.UpdateEvent(eventID, &updateEventRequest, userID)
	if eventErr != nil {
		ctx.AbortWithStatusJSON(eventErr.Code, eventErr)
		return
	}

	ctx.JSON(200, event)
}

func (h *Event) DeleteEvent(ctx *gin.Context) {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Event ID is not valid"})
		return
	}

	userID := helpers.GetUserId(ctx)

	event, eventErr := h.eventService.DeleteEvent(eventID, userID)
	if eventErr != nil {
		ctx.AbortWithStatusJSON(eventErr.Code, eventErr)
		return
	}

	ctx.JSON(200, event)
}
