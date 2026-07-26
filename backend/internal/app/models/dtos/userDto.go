package dtos

import (
	"backend/internal/app/entities"
)

type UserResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type GetAllUsersResponse struct {
	Users   []*UserResponse `json:"users"`
	Message string          `json:"message"`
}

type CreateUserRequest struct {
	Email       string `json:"email" binding:"required,email,max=254"`
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Password    string `json:"password" binding:"required,min=8,max=100"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type UpdateUserRequest struct {
	Email       string `json:"email" binding:"omitempty,email,max=254"`
	Name        string `json:"name" binding:"omitempty,min=3,max=100"`
	Password    string `json:"password" binding:"omitempty,min=8,max=100"`
	PhoneNumber string `json:"phone_number" binding:"omitempty"`
}

type CreateUserResponse struct {
	Email       string `json:"email" binding:"required,email,max=254"`
	Name        string `json:"name" binding:"required,min=3,max=100"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Message     string `json:"message" binding:"required"`
}

func (r *GetAllUsersResponse) MapUsersResponse(users []*entities.User) {
	for _, user := range users {
		user := &UserResponse{
			Email:       user.Email,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		}
		r.Users = append(r.Users, user)
	}
}

func (r *UserResponse) MapUserResponse(user *entities.User) {
	r.Email = user.Email
	r.Name = user.Name
	r.PhoneNumber = user.PhoneNumber
}

func (ur *CreateUserRequest) ToUser() *entities.User {

	return &entities.User{
		Email:       ur.Email,
		Name:        ur.Name,
		Password:    ur.Password,
		PhoneNumber: ur.PhoneNumber,
	}
}

func (ur *UpdateUserRequest) ToUser() *entities.User {
	return &entities.User{
		Email:       ur.Email,
		Name:        ur.Name,
		Password:    ur.Password,
		PhoneNumber: ur.PhoneNumber,
	}
}

func (ur *CreateUserResponse) FromUser(user *entities.User) *CreateUserResponse {
	return &CreateUserResponse{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Message:     "User created successfully",
	}
}
