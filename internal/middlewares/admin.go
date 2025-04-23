package middlewares

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/gin-gonic/gin"
)

func NewAdminMiddleware(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC401_CODE,
				Message: "User ID not found in context",
				From:    "AdminMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, exceptions.ErrorMessage{
					Title:  "Unauthorized",
					Detail: "User ID not found in context.",
				}),
			})
			return
		}

		user, err := userRepo.GetUser(userID.(string))
		if err != nil || !user.IsAdmin {
			logging.NewLogger(logging.Logger{
				Code:    exceptions.RFC403_CODE,
				Message: "User is not an administrator",
				From:    "AdminMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Forbidden, exceptions.ErrorMessage{
					Title:  "Forbidden",
					Detail: "User is not an administrator.",
				}),
			})
			return
		}

		c.Next()
	}
}
