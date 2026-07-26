package handlers

import (
	"backend/internal/app/helpers"
	"backend/internal/app/models/dtos"
	"backend/internal/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Todo struct {
	todoService *services.Todo
	logger      zerolog.Logger
}

func NewTodoHandler(todoService *services.Todo, logger zerolog.Logger) *Todo {
	return &Todo{todoService: todoService, logger: logger}
}

func (h *Todo) GetAllTodos(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)

	allTodos, err := h.todoService.GetAllTodos(userID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}

	h.logger.Debug().Any("allTodos", allTodos).Msg("Retrieved all todos successfully")

	ctx.JSON(200, allTodos)
}

func (h *Todo) GetTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	userID := helpers.GetUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Todo ID is not valid"})
		return
	}

	todo, todoErr := h.todoService.GetTodo(todoID, userID)
	if todoErr != nil {
		ctx.AbortWithStatusJSON(todoErr.Code, todoErr)
		return
	}

	ctx.JSON(200, todo)
}

func (h *Todo) CreateTodo(ctx *gin.Context) {
	userID := helpers.GetUserId(ctx)
	var createTodoRequest dtos.CreateTodoRequest
	if err := ctx.ShouldBindJSON(&createTodoRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	todo, todoErr := h.todoService.CreateTodo(&createTodoRequest, userID)
	if todoErr != nil {
		ctx.AbortWithStatusJSON(todoErr.Code, todoErr)
		return
	}

	ctx.JSON(201, todo)
}

func (h *Todo) UpdateTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	userID := helpers.GetUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Todo ID is not valid"})
		return
	}

	var updateTodoRequest dtos.UpdateTodoRequest
	if err := ctx.ShouldBindJSON(&updateTodoRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	todo, todoErr := h.todoService.UpdateTodo(todoID, &updateTodoRequest, userID)
	if todoErr != nil {
		ctx.AbortWithStatusJSON(todoErr.Code, todoErr)
		return
	}

	ctx.JSON(200, todo)
}

func (h *Todo) UpdateCompletedStatus(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	userID := helpers.GetUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Todo ID is not valid"})
		return
	}

	todo, todoErr := h.todoService.UpdateCompletedStatus(todoID, userID)
	if todoErr != nil {
		ctx.AbortWithStatusJSON(todoErr.Code, todoErr)
		return
	}

	ctx.JSON(200, todo)
}

func (h *Todo) UpdateInCompletedStatus(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	userID := helpers.GetUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Todo ID is not valid"})
		return
	}

	todo, todoErr := h.todoService.UpdateIncompleteStatus(todoID, userID)
	if todoErr != nil {
		ctx.AbortWithStatusJSON(todoErr.Code, todoErr)
		return
	}

	ctx.JSON(200, todo)
}

func (h *Todo) DeleteTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	userID := helpers.GetUserId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Todo ID is not valid"})
		return
	}

	todo, todoErr := h.todoService.DeleteTodo(todoID, userID)
	if todoErr != nil {
		ctx.AbortWithStatusJSON(todoErr.Code, todoErr)
		return
	}

	ctx.JSON(200, todo)
}
