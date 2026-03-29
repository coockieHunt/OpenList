package handler

import (
	"OpenList/Go/service/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

// reponse wrapper for consistent API responses
func SendResponse(c *gin.Context, status string, message string, data interface{}) {
	Response := sqlite.APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.IndentedJSON(http.StatusOK, Response)
}
