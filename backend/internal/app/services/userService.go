package services

import (
	"backend/internal/app/helpers"
	"backend/internal/app/models"
	"backend/internal/app/models/dtos"
	"backend/internal/app/repositories"
	"database/sql"
	"errors"
	"net/http"
)

type User struct {
	userRepository *repositories.User
}

func NewUserService(userRepository *repositories.User) *User {
	return &User{userRepository: userRepository}
}

func (us *User) GetAllUsers() (*dtos.GetAllUsersResponse, *models.ErrorResponse) {
	response := &dtos.GetAllUsersResponse{}

	queriedUsers, err := us.userRepository.GetAllUsers()
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve users",
		}
	}

	response.MapUsersResponse(queriedUsers)
	return response, nil
}

func (us *User) GetUser(userID int) (*dtos.UserResponse, *models.ErrorResponse) {
	response := &dtos.UserResponse{}

	user, err := us.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user",
		}
	}

	response.MapUserResponse(user)

	return response, nil
}

func (us *User) DeleteUser(userID int) *models.ErrorResponse {
	_, err := us.userRepository.FindByID(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}

		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user",
		}
	}

	err = us.userRepository.DeleteUser(userID)
	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to delete user",
		}
	}

	return nil
}

func (us *User) CreateUser(createUserRequest *dtos.CreateUserRequest) (*dtos.CreateUserResponse, *models.ErrorResponse) {
	userResponse := &dtos.CreateUserResponse{}

	errEmail := us.checkIfEmailExists(createUserRequest.Email)
	if errEmail != nil {
		return nil, errEmail
	}

	hashedPassword, hashErr := helpers.HashPassword(createUserRequest.Password)
	if hashErr != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to hash password",
		}
	}

	user := createUserRequest.ToUser()
	user.Password = hashedPassword

	err := us.userRepository.CreateUser(user)
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user",
		}
	}

	return userResponse.FromUser(user), nil
}

func (us *User) UpdateUser(userID int, updateUserRequest *dtos.UpdateUserRequest) *models.ErrorResponse {
	existingUser, err := us.userRepository.FindByIDWithPassword(userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "User not found",
			}
		}

		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve user",
		}
	}

	if updateUserRequest.Email != "" {
		if updateUserRequest.Email != existingUser.Email {
			errEmail := us.checkIfEmailExists(updateUserRequest.Email)
			if errEmail != nil {
				return errEmail
			}
			existingUser.Email = updateUserRequest.Email
		}
	}

	if updateUserRequest.Password != "" {
		hashedPassword, hashErr := helpers.HashPassword(updateUserRequest.Password)
		if hashErr != nil {
			return &models.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "Failed to hash password",
			}
		}
		existingUser.Password = hashedPassword
	}

	if updateUserRequest.Name != "" {
		existingUser.Name = updateUserRequest.Name
	}

	if updateUserRequest.PhoneNumber != "" {
		if updateUserRequest.PhoneNumber != existingUser.PhoneNumber {
			errPhone := us.checkIfPhoneNumberExists(updateUserRequest.PhoneNumber)
			if errPhone != nil {
				return errPhone
			}
			existingUser.PhoneNumber = updateUserRequest.PhoneNumber
		}
	}

	existingUser.ID = userID

	err = us.userRepository.UpdateUser(existingUser)

	if err != nil {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update user",
		}
	}

	return nil

}

func (us *User) checkIfEmailExists(email string) *models.ErrorResponse {
	userWithEmail, err := us.userRepository.FindByEmail(email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to check email existence",
		}
	}

	if userWithEmail != nil {
		return &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Email already exists",
		}
	}
	return nil
}

func (us *User) checkIfPhoneNumberExists(phoneNumber string) *models.ErrorResponse {
	userWithPhoneNumber, err := us.userRepository.FindByPhoneNumber(phoneNumber)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to check phone number existence",
		}
	}

	if userWithPhoneNumber != nil {
		return &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Phone number already exists",
		}
	}
	return nil
}
