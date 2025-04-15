package middlewares

import (
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

func NewAdminMiddleware(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			util.NewLogger(util.Logger{
				Code:    util.RFC401_CODE,
				Message: "User ID not found in context",
				From:    "AdminMiddleware",
				Layer:   "Infra",
				TypeLog: "ERROR",
			})

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": util.NewProblemDetails(util.Unauthorized, util.ErrorMessage{
					Title:  "Unauthorized",
					Detail: "User ID not found in context.",
				}),
			})
			return
		}

		user, err := userRepo.GetUser(userID.(string))
		if err != nil || !user.IsAdmin {
			util.NewLogger(util.Logger{
				Code:    util.RFC403_CODE,
				Message: "User is not an administrator",
				From:    "AdminMiddleware",
				Layer:   "Infra",
				TypeLog: "ERROR",
			})

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": util.NewProblemDetails(util.Forbidden, util.ErrorMessage{
					Title:  "Forbidden",
					Detail: "User is not an administrator.",
				}),
			})
			return
		}

		c.Next()
	}
}
