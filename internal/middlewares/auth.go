package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			util.NewLogger(util.Logger{
				Code:    401,
				Message: util.GetErrorMessage("AuthMiddleware", "UnauthorizedHeader").Detail,
				From:    "AuthMiddleware",
				Layer:   util.LoggerLayers.MIDDLEWARES,
				TypeLog: util.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, util.GetErrorMessage("AuthMiddleware", "UnauthorizedHeader")),
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			util.NewLogger(util.Logger{
				Code:    401,
				Message: util.GetErrorMessage("AuthMiddleware", "UnauthorizedBearer").Detail,
				From:    "AuthMiddleware",
				Layer:   util.LoggerLayers.MIDDLEWARES,
				TypeLog: util.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, util.GetErrorMessage("AuthMiddleware", "UnauthorizedBearer")),
			})
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				util.NewLogger(util.Logger{
					Code:    401,
					Message: util.GetErrorMessage("AuthMiddleware", "UnauthorizedTokenParse").Detail,
					From:    "AuthMiddleware",
					Layer:   util.LoggerLayers.MIDDLEWARES,
					TypeLog: util.LoggerTypes.ERROR,
				})
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": exceptions.NewProblemDetails(exceptions.Unauthorized, util.GetErrorMessage("AuthMiddleware", "UnauthorizedTokenParse")),
				})
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(config.SECRETS_VAR.JWT_SECRET), nil
		})

		if err != nil {
			util.NewLogger(util.Logger{
				Code:    401,
				Message: util.GetErrorMessage("AuthMiddleware", "UnauthorizedInvalidToken").Detail,
				From:    "AuthMiddleware",
				Layer:   util.LoggerLayers.MIDDLEWARES,
				TypeLog: util.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, util.GetErrorMessage("AuthMiddleware", "UnauthorizedInvalidToken")),
			})
			return
		}

		if !token.Valid {
			util.NewLogger(util.Logger{
				Code:    401,
				Message: util.GetErrorMessage("AuthMiddleware", "UnauthorizedToken").Detail,
				From:    "AuthMiddleware",
				Layer:   util.LoggerLayers.MIDDLEWARES,
				TypeLog: util.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, util.GetErrorMessage("AuthMiddleware", "UnauthorizedToken")),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)

		user, err := userRepo.GetUser(userID)
		if err != nil {
			util.NewLogger(util.Logger{
				Code:    401,
				Message: util.GetErrorMessage("LoginUseCase", "UserNotFound").Detail,
				From:    "AuthMiddleware",
				Layer:   util.LoggerLayers.MIDDLEWARES,
				TypeLog: util.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.NotFound, util.GetErrorMessage("LoginUseCase", "UserNotFound")),
			})
			return
		}

		if !user.Active {
			util.NewLogger(util.Logger{
				Code:    401,
				Message: util.GetErrorMessage("LoginUseCase", "UserNotActive").Detail,
				From:    "AuthMiddleware",
				Layer:   util.LoggerLayers.MIDDLEWARES,
				TypeLog: util.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Forbidden, util.GetErrorMessage("LoginUseCase", "UserNotActive")),
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
