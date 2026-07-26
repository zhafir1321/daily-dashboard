package router

import (
	"backend/internal/app/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPublicEndpoints(router *gin.Engine, userHandlers *handlers.User, middleware *middleware.Middleware) {
	api := router.Group("/api").Use(middleware.Authentication())
	{
		api.GET("/users", userHandlers.GetAllUsers)
		api.GET("/users/:id", userHandlers.GetUser)
		api.POST("/users", userHandlers.CreateUser)
		api.PUT("/users/:id", userHandlers.UpdateUser)
		api.DELETE("/users/:id", userHandlers.DeleteUser)
	}
}

func RegisterAuthEndpoints(router *gin.Engine, authHandlers *handlers.Auth) {
	api := router.Group("/auth")
	{
		api.POST("/register", authHandlers.RegisterUser)
		api.POST("/login", authHandlers.LoginUser)
	}
}

func RegisterTodoEndpoints(router *gin.Engine, todoHandlers *handlers.Todo, middleware *middleware.Middleware) {
	api := router.Group("/api").Use(middleware.Authentication())
	{
		api.GET("/todos", todoHandlers.GetAllTodos)
		api.GET("/todos/:id", todoHandlers.GetTodo)
		api.POST("/todos", todoHandlers.CreateTodo)
		api.PUT("/todos/:id", todoHandlers.UpdateTodo)
		api.PATCH("/todos/:id/complete", todoHandlers.UpdateCompletedStatus)
		api.PATCH("/todos/:id/incomplete", todoHandlers.UpdateInCompletedStatus)
		api.DELETE("/todos/:id", todoHandlers.DeleteTodo)
	}
}
