package auth

import (
	"OpenList/Go/service/sqlite"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// EnsureDefaultUser: Creates a default user if no users exist, returns true if created along with credentials
func EnsureDefaultUser() (bool, string, string, error) {
	var count int64
	if err := sqlite.DB.Model(&sqlite.User{}).Count(&count).Error; err != nil {
		return false, "", "", err
	}
	if count > 0 {
		return false, "", "", nil
	}

	username := os.Getenv("OPENLIST_DEFAULT_USER")
	if username == "" {
		username = "admin"
	}

	password := os.Getenv("OPENLIST_DEFAULT_PASSWORD")
	if password == "" {
		password = generatePassword(12)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, "", "", err
	}

	user := sqlite.User{
		Username:     username,
		PasswordHash: string(hash),
		FirstLogin:   true,
	}

	if err := sqlite.DB.Create(&user).Error; err != nil {
		return false, "", "", err
	}

	return true, username, password, nil
}

// Authenticate: Verifies username and password, returns user if valid
func Authenticate(username, password string) (*sqlite.User, error) {
	var user sqlite.User
	if err := sqlite.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return &user, nil
}

// ChangePassword: Allows users to change their password
func ChangePassword(userID uint, currentPassword, newPassword string) error {
	if len(newPassword) < 8 {
		return errors.New("new password must be at least 8 characters")
	}

	var user sqlite.User
	if err := sqlite.DB.First(&user, userID).Error; err != nil {
		return ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(currentPassword)); err != nil {
		return ErrInvalidCredentials
	}

	return setPassword(&user, newPassword)
}

// SetSingleUserPassword: Sets the password for the single user in single-user mode (used for CLI password setup)
func SetSingleUserPassword(newPassword string) error {
	if len(newPassword) < 8 {
		return errors.New("new password must be at least 8 characters")
	}

	var user sqlite.User
	if err := sqlite.DB.Order("id asc").First(&user).Error; err != nil {
		return err
	}

	return setPassword(&user, newPassword)
}

// Internal function to update a user's password and invalidate sessions
func setPassword(user *sqlite.User, newPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)
	user.FirstLogin = false
	if err := sqlite.DB.Save(user).Error; err != nil {
		return err
	}

	if err := sqlite.DB.Where("user_id = ?", user.ID).Delete(&sqlite.Session{}).Error; err != nil {
		return err
	}

	return nil
}

// CreateSession: Generates a new session token for a user
func CreateSession(userID uint) (string, error) {
	token := generateToken()
	session := sqlite.Session{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err := sqlite.DB.Create(&session).Error; err != nil {
		return "", err
	}

	return token, nil
}

// ValidateSession: Validates a session token and returns the associated user
func ValidateSession(token string) (*sqlite.User, error) {
	if token == "" {
		return nil, ErrInvalidCredentials
	}

	var session sqlite.Session
	if err := sqlite.DB.Where("token = ?", token).First(&session).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	if time.Now().After(session.ExpiresAt) {
		_ = sqlite.DB.Delete(&session).Error
		return nil, ErrInvalidCredentials
	}

	var user sqlite.User
	if err := sqlite.DB.First(&user, session.UserID).Error; err != nil {
		return nil, ErrInvalidCredentials
	}

	return &user, nil
}

// DeleteSession: Deletes a session token (used for logout)
func DeleteSession(token string) {
	if token == "" {
		return
	}

	_ = sqlite.DB.Where("token = ?", token).Delete(&sqlite.Session{}).Error
}

// Helper functions
func generateToken() string {
	bytes := make([]byte, 24)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}

// generatePassword: Creates a random password of specified length (used for default user)
func generatePassword(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "ChangeMe123!"
	}

	return hex.EncodeToString(bytes)[:length]
}
