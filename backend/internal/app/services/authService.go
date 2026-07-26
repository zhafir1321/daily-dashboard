package services

import (
	"backend/internal/app/entities"
	"backend/internal/app/helpers"
	"backend/internal/app/models"
	"backend/internal/app/models/dtos"
	"backend/internal/app/repositories"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type Auth struct {
	userRepository *repositories.User
	JWTSecret      string
}

func NewAuthService(userRepository *repositories.User, jwtSecret string) *Auth {
	return &Auth{
		userRepository: userRepository,
		JWTSecret:      jwtSecret,
	}
}

func (as *Auth) LoginUser(loginReq *dtos.LoginRequest) (*dtos.LoginResponse, *models.ErrorResponse) {
	response := &dtos.LoginResponse{}

	user, userErr := as.userRepository.FindByEmailWithPassword(loginReq.Email)
	if userErr != nil {
		if errors.Is(userErr, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid email or password",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user",
		}
	}

	if !helpers.CheckPasswordHash(loginReq.Password, user.Password) {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid email or password",
		}
	}

	claims := helpers.Claims{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	token, tokenErr := helpers.CreateToken(claims, as.JWTSecret)
	if tokenErr != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		}
	}

	response.MapLoginResponse(*token)

	return response, nil
}

func (as *Auth) RegisterUser(registerReq *dtos.RegisterRequest) (*dtos.RegisterResponse, *models.ErrorResponse) {
	response := &dtos.RegisterResponse{}

	userWithEmail, emailErr := as.userRepository.FindByEmail(registerReq.Email)
	if emailErr != nil && !errors.Is(emailErr, sql.ErrNoRows) {
		fmt.Println("Error checking email:", emailErr)
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to check email",
		}
	}
	if userWithEmail != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already exists",
		}
	}

	userWithPhone, phoneErr := as.userRepository.FindByPhoneNumber(registerReq.PhoneNumber)
	if phoneErr != nil && !errors.Is(phoneErr, sql.ErrNoRows) {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to check phone number",
		}
	}
	if userWithPhone != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Phone number already exists",
		}
	}

	hashedPassword, hashErr := helpers.HashPassword(registerReq.Password)
	if hashErr != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to hash password",
		}
	}

	newUser := &entities.User{
		Email:       registerReq.Email,
		Name:        registerReq.Name,
		Password:    hashedPassword,
		PhoneNumber: registerReq.PhoneNumber,
	}

	createErr := as.userRepository.CreateUser(newUser)
	if createErr != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
		}
	}

	response.MapRegisterResponse(newUser.Email, newUser.Name, newUser.PhoneNumber)

	return response, nil
}
