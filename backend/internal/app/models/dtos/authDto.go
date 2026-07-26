package dtos

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=254"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email       string `json:"email" binding:"required,email,max=254"`
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Password    string `json:"password" binding:"required,min=8,max=100"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type RegisterResponse struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Message     string `json:"message"`
}

func (lr *LoginResponse) MapLoginResponse(token string) {
	lr.Token = token
	lr.Message = "Login successful"
}

func (rr *RegisterResponse) MapRegisterResponse(email, name, phoneNumber string) {
	rr.Email = email
	rr.Name = name
	rr.PhoneNumber = phoneNumber
	rr.Message = "User registered successfully"
}
