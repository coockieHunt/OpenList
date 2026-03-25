package routes

import (
	"OpenList/Go/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GenrateAuthToken(c *gin.Context) {
	token, err := auth.GenerateAuthToken(1, false)
	if err == nil && token != "" {
		fmt.Print(token)
		SendResponse(c, "success", "Token generated", token)
	} else {
		SendResponse(c, "error", "Error generating token", nil)
	}
}
