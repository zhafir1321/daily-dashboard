package handlers

import (
	"backend/internal/app/models/dtos"
	"backend/internal/app/services"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	authService *services.Auth
}

func NewAuthHandler(authService *services.Auth) *Auth {
	return &Auth{authService: authService}
}

func (h *Auth) LoginUser(ctx *gin.Context) {
	var loginRequest dtos.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	loginResponse, loginErr := h.authService.LoginUser(&loginRequest)
	if loginErr != nil {
		ctx.AbortWithStatusJSON(loginErr.Code, loginErr)
		return
	}

	ctx.JSON(200, loginResponse)

}

func (h *Auth) RegisterUser(ctx *gin.Context) {
	var registerRequest dtos.RegisterRequest

	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	registerResponse, registerErr := h.authService.RegisterUser(&registerRequest)
	if registerErr != nil {
		ctx.AbortWithStatusJSON(registerErr.Code, registerErr)
		return
	}

	ctx.JSON(201, registerResponse)
}
