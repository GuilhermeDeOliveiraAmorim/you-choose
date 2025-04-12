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
