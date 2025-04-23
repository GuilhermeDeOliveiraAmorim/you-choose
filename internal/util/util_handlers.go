package util

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/gin-gonic/gin"
)

func HandleErrors(c *gin.Context, errs []exceptions.ProblemDetails) {
	if len(errs) == 0 {
		return
	}

	status := errs[0].Status
	c.JSON(status, gin.H{"errors": errs})
}

func GetUserID(c *gin.Context) (string, *exceptions.ProblemDetails) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", &exceptions.ProblemDetails{
			Type:     "Unauthorized",
			Title:    "Missing User ID",
			Status:   http.StatusUnauthorized,
			Detail:   "User id is required",
			Instance: exceptions.RFC401,
		}
	}

	userIDStr, ok := userID.(string)
	if !ok || userIDStr == "" {
		return "", &exceptions.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Invalid User ID",
			Status:   http.StatusBadRequest,
			Detail:   "A valid user id is required",
			Instance: exceptions.RFC400,
		}
	}

	return userIDStr, nil
}
