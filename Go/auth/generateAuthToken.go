package auth

import (
	"OpenList/Go/sqlite"
	"OpenList/Go/token"
	"os"
	"strconv"
	"time"
)

func GenerateAuthToken(userID uint, admin bool) (string, error) {
	newToken := token.GenToken()
	curDate := time.Now()
	expirationDuration := os.Getenv("TOKEN_EXPIRATION_HOURS")

	if expirationDuration == "" {
		expirationDuration = "24"
	}

	expirationDurationFormat, _ := strconv.Atoi(expirationDuration)
	ExpiresAt := curDate.Add(time.Duration(expirationDurationFormat) * time.Hour)

	// fmt.Print(ExpiresAt, curDate, expirationDurationFormat)

	authToken := sqlite.AuthToken{
		UserID:    userID,
		Token:     newToken,
		Admin:     admin,
		ExpiresAt: ExpiresAt,
	}

	if err := sqlite.DB.Create(&authToken).Error; err != nil {
		return "", err
	}
	return newToken, nil
}
