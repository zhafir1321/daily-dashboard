package middleware

import (
	"backend/internal/app/helpers"
	"backend/internal/app/repositories"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	JWTSecret      string
	userRepository *repositories.User
}

func NewMiddleware(userRepository *repositories.User, jwtSecret string) *Middleware {
	return &Middleware{
		userRepository: userRepository,
		JWTSecret:      jwtSecret,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if bearer == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is missing"})
			return
		}

		if !strings.HasPrefix(bearer, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := strings.TrimPrefix(bearer, "Bearer ")

		claims, err := helpers.ParseToken(tokenString, m.JWTSecret)
		if err != nil {
			c.AbortWithStatusJSON(err.Code, gin.H{"error": err.Message})
			return
		}

		c.Set("id", claims.ID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)
		c.Set("phone_number", claims.PhoneNumber)

		c.Next()
	}
}
