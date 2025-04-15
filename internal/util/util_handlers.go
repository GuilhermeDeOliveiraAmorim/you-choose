package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrors(c *gin.Context, errs []ProblemDetails) {
	if len(errs) == 0 {
		return
	}

	status := errs[0].Status
	c.JSON(status, gin.H{"errors": errs})
}

func GetUserID(c *gin.Context) (string, *ProblemDetails) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", &ProblemDetails{
			Type:     "Unauthorized",
			Title:    "Missing User ID",
			Status:   http.StatusUnauthorized,
			Detail:   "User id is required",
			Instance: RFC401,
		}
	}

	userIDStr, ok := userID.(string)
	if !ok || userIDStr == "" {
		return "", &ProblemDetails{
			Type:     "Bad Request",
			Title:    "Invalid User ID",
			Status:   http.StatusBadRequest,
			Detail:   "A valid user id is required",
			Instance: RFC400,
		}
	}

	return userIDStr, nil
}
