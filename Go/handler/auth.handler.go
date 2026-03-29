package handler

import (
	auth "OpenList/Go/service/auth"
	"OpenList/Go/service/sqlite"
	"strings"

	"github.com/gin-gonic/gin"
)

const sessionCookieName = "openlist_session"

// Login: Authenticates a user and creates a session
func Login(c *gin.Context) {
	var payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		SendResponse(c, "error", "invalid payload", nil)
		return
	}

	user, err := auth.Authenticate(payload.Username, payload.Password)
	if err != nil {
		SendResponse(c, "error", "invalid credentials", nil)
		return
	}

	token, err := auth.CreateSession(user.ID)
	if err != nil {
		SendResponse(c, "error", "failed to create session", nil)
		return
	}

	c.SetCookie(sessionCookieName, token, 24*3600, "/", "", false, true)
	SendResponse(c, "success", "login successful", gin.H{
		"first_login": user.FirstLogin,
	})
}

// Logout: Deletes the user's session
func Logout(c *gin.Context) {
	token := readSessionToken(c)
	auth.DeleteSession(token)
	c.SetCookie(sessionCookieName, "", -1, "/", "", false, true)
	SendResponse(c, "success", "logout successful", nil)
}

// ChangePassword: Allows users to change their password
func ChangePassword(c *gin.Context) {
	userID := c.GetUint("userID")

	var payload struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		SendResponse(c, "error", "invalid payload", nil)
		return
	}

	if err := auth.ChangePassword(userID, payload.CurrentPassword, payload.NewPassword); err != nil {
		SendResponse(c, "error", err.Error(), nil)
		return
	}

	c.SetCookie(sessionCookieName, "", -1, "/", "", false, true)
	SendResponse(c, "success", "password updated", nil)
}

func AuthStatus(c *gin.Context) {
	userValue, exists := c.Get("user")
	if !exists {
		SendResponse(c, "error", "unauthorized", nil)
		return
	}

	user := userValue.(*sqlite.User)
	SendResponse(c, "success", "session valid", gin.H{
		"username":    user.Username,
		"first_login": user.FirstLogin,
	})
}

func readSessionToken(c *gin.Context) string {
	if cookie, err := c.Cookie(sessionCookieName); err == nil && cookie != "" {
		return cookie
	}

	header := c.GetHeader("Authorization")
	if strings.HasPrefix(header, "Bearer ") {
		return strings.TrimPrefix(header, "Bearer ")
	}

	return ""
}
