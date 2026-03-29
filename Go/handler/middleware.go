package handler

import (
	auth "OpenList/Go/service/auth"
	"OpenList/Go/service/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRequired: Middleware to protect routes that require authentication
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := readSessionToken(c)
		user, err := auth.ValidateSession(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "unauthorized",
			})
			return
		}

		c.Set("userID", user.ID)
		c.Set("user", user)
		c.Next()
	}
}

// MustChangePasswordGuard: Middleware to enforce password change on first login
func MustChangePasswordGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		userValue, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "unauthorized"})
			return
		}

		user := userValue.(*sqlite.User)
		if user.FirstLogin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "password change required",
			})
			return
		}

		c.Next()
	}
}
