package routes

import (
	"OpenList/Go/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, status string, message string, data interface{}) {
	Response := sqlite.APIResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
	c.IndentedJSON(http.StatusOK, Response)
}
