package handlers

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

func handleErrors(c *gin.Context, errs []util.ProblemDetails) {
	if len(errs) > 0 {
		for _, err := range errs {
			if err.Status == 500 {
				c.JSON(err.Status, gin.H{"error": err})
				return
			} else {
				c.JSON(err.Status, gin.H{"error": err})
				return
			}
		}
	}
}

func getUserID(c *gin.Context) (string, *util.ProblemDetails) {
	userID, exists := c.Get("userID")
	if !exists {
		return "", &util.ProblemDetails{
			Type:     "Unauthorized",
			Title:    "Missing User ID",
			Status:   http.StatusUnauthorized,
			Detail:   "User id is required",
			Instance: util.RFC401,
		}
	}

	userIDStr, ok := userID.(string)
	if !ok || userIDStr == "" {
		return "", &util.ProblemDetails{
			Type:     "Bad Request",
			Title:    "Invalid User ID",
			Status:   http.StatusBadRequest,
			Detail:   "A valid user id is required",
			Instance: util.RFC400,
		}
	}

	return userIDStr, nil
}
