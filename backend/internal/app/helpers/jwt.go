package helpers

import (
	"backend/internal/app/models"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func CreateToken(c Claims, secret string) (*string, *models.ErrorResponse) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           c.ID,
		"name":         c.Name,
		"email":        c.Email,
		"phone_number": c.PhoneNumber,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtSecret := []byte(secret)

	tokenString, tokenErr := token.SignedString(jwtSecret)
	if tokenErr != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create token",
		}
	}
	return &tokenString, nil
}

func ParseToken(tokenString string, secret string) (*Claims, *models.ErrorResponse) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, &models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		}
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, &models.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
}
