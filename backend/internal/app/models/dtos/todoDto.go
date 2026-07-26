package dtos

import "backend/internal/app/entities"

type TodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetTodoResponse struct {
	Todo    *TodoResponse `json:"todo"`
	Message string        `json:"message"`
}

type GetAllTodosResponse struct {
	Todos   []*TodoResponse `json:"todos"`
	Message string          `json:"message"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
	Completed   *bool  `json:"completed" binding:"omitempty"`
}

type CreateTodoResponse struct {
	Message string `json:"message"`
}

type DeleteTodoResponse struct {
	Message string `json:"message"`
}

func (r *GetAllTodosResponse) MapTodosResponse(todos []*entities.Todo) {
	for _, todo := range todos {
		todoResponse := &TodoResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
		r.Todos = append(r.Todos, todoResponse)
	}
}

func (r *TodoResponse) MapTodoResponse(todo *entities.Todo) {
	r.ID = todo.ID
	r.Title = todo.Title
	r.Description = todo.Description
	r.Completed = todo.Completed
	r.CreatedAt = todo.CreatedAt
	r.UpdatedAt = todo.UpdatedAt
}

func (tr *CreateTodoRequest) ToTodo() *entities.Todo {
	return &entities.Todo{
		Title:       tr.Title,
		Description: tr.Description,
		Completed:   false,
	}
}
