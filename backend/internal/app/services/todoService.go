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

type Todo struct {
	todoRepository *repositories.Todo
	logger         zerolog.Logger
}

func NewTodoService(todoRepository *repositories.Todo, logger zerolog.Logger) *Todo {
	return &Todo{todoRepository: todoRepository, logger: logger}
}

func (ts *Todo) GetAllTodos(userID int) (*dtos.GetAllTodosResponse, *models.ErrorResponse) {
	response := &dtos.GetAllTodosResponse{}

	queriedTodos, err := ts.todoRepository.GetAllTodos(userID)
	if err != nil {
		ts.logger.Error().Err(err).Msg("Failed to retrieve todos")
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todos",
		}
	}

	ts.logger.Debug().Any("todos", queriedTodos).Msg("this is response")
	response.MapTodosResponse(queriedTodos)
	response.Message = "Todos retrieved successfully"
	return response, nil
}

func (ts *Todo) GetTodo(todoID int, userID int) (*dtos.GetTodoResponse, *models.ErrorResponse) {
	response := &dtos.GetTodoResponse{}
	response.Todo = &dtos.TodoResponse{}

	queriedTodo, queryErr := ts.todoRepository.FindByID(todoID, userID)
	if queryErr != nil {
		if errors.Is(queryErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Todo not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo",
		}
	}

	response.Todo.MapTodoResponse(queriedTodo)
	response.Message = "Todo retrieved successfully"

	return response, nil
}

func (ts *Todo) CreateTodo(todoReq *dtos.CreateTodoRequest, userID int) (*dtos.CreateTodoResponse, *models.ErrorResponse) {
	todo := todoReq.ToTodo()
	todo.UserID = userID
	err := ts.todoRepository.CreateTodo(todo)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create todo",
		}
	}

	response := &dtos.CreateTodoResponse{
		Message: "Todo created successfully",
	}

	return response, nil
}

func (ts *Todo) UpdateTodo(todoID int, todoReq *dtos.UpdateTodoRequest, userID int) (*dtos.CreateTodoResponse, *models.ErrorResponse) {

	existingTodo, err := ts.todoRepository.FindByID(todoID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Todo not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo",
		}
	}

	if todoReq.Title != "" {
		existingTodo.Title = todoReq.Title
	}

	if todoReq.Description != "" {
		existingTodo.Description = todoReq.Description
	}

	err = ts.todoRepository.UpdateTodo(existingTodo)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update todo",
		}
	}

	response := &dtos.CreateTodoResponse{
		Message: "Todo updated successfully",
	}

	return response, nil

}

func (ts *Todo) UpdateCompletedStatus(todoID int, userID int) (*dtos.CreateTodoResponse, *models.ErrorResponse) {
	existingTodo, err := ts.todoRepository.FindByID(todoID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Todo not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo",
		}
	}

	existingTodo.Completed = true
	err = ts.todoRepository.UpdateTodo(existingTodo)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update todo",
		}
	}

	response := &dtos.CreateTodoResponse{
		Message: "Todo updated successfully",
	}

	return response, nil
}

func (ts *Todo) UpdateIncompleteStatus(todoID int, userID int) (*dtos.CreateTodoResponse, *models.ErrorResponse) {
	existingTodo, err := ts.todoRepository.FindByID(todoID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Todo not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo",
		}
	}

	existingTodo.Completed = false
	err = ts.todoRepository.UpdateTodo(existingTodo)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update todo",
		}
	}

	response := &dtos.CreateTodoResponse{
		Message: "Todo updated successfully",
	}

	return response, nil
}

func (ts *Todo) DeleteTodo(todoID int, userID int) (*dtos.DeleteTodoResponse, *models.ErrorResponse) {
	existingTodo, err := ts.todoRepository.FindByID(todoID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Todo not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve todo",
		}
	}

	err = ts.todoRepository.DeleteTodo(existingTodo.ID, userID)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete todo",
		}
	}

	response := &dtos.DeleteTodoResponse{
		Message: "Todo deleted successfully",
	}

	return response, nil
}
