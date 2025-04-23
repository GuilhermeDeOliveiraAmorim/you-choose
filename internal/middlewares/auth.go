package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			logging.NewLogger(logging.Logger{
				Code:    401,
				Message: language.GetErrorMessage("AuthMiddleware", "UnauthorizedHeader").Detail,
				From:    "AuthMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("AuthMiddleware", "UnauthorizedHeader")),
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logging.NewLogger(logging.Logger{
				Code:    401,
				Message: language.GetErrorMessage("AuthMiddleware", "UnauthorizedBearer").Detail,
				From:    "AuthMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("AuthMiddleware", "UnauthorizedBearer")),
			})
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logging.NewLogger(logging.Logger{
					Code:    401,
					Message: language.GetErrorMessage("AuthMiddleware", "UnauthorizedTokenParse").Detail,
					From:    "AuthMiddleware",
					Layer:   logging.LoggerLayers.MIDDLEWARES,
					TypeLog: logging.LoggerTypes.ERROR,
				})
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("AuthMiddleware", "UnauthorizedTokenParse")),
				})
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(config.SECRETS_VAR.JWT_SECRET), nil
		})

		if err != nil {
			logging.NewLogger(logging.Logger{
				Code:    401,
				Message: language.GetErrorMessage("AuthMiddleware", "UnauthorizedInvalidToken").Detail,
				From:    "AuthMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("AuthMiddleware", "UnauthorizedInvalidToken")),
			})
			return
		}

		if !token.Valid {
			logging.NewLogger(logging.Logger{
				Code:    401,
				Message: language.GetErrorMessage("AuthMiddleware", "UnauthorizedToken").Detail,
				From:    "AuthMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Unauthorized, language.GetErrorMessage("AuthMiddleware", "UnauthorizedToken")),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)

		user, err := userRepo.GetUser(userID)
		if err != nil {
			logging.NewLogger(logging.Logger{
				Code:    401,
				Message: language.GetErrorMessage("LoginUseCase", "UserNotFound").Detail,
				From:    "AuthMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.NotFound, language.GetErrorMessage("LoginUseCase", "UserNotFound")),
			})
			return
		}

		if !user.Active {
			logging.NewLogger(logging.Logger{
				Code:    401,
				Message: language.GetErrorMessage("LoginUseCase", "UserNotActive").Detail,
				From:    "AuthMiddleware",
				Layer:   logging.LoggerLayers.MIDDLEWARES,
				TypeLog: logging.LoggerTypes.ERROR,
			})
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": exceptions.NewProblemDetails(exceptions.Forbidden, language.GetErrorMessage("LoginUseCase", "UserNotActive")),
			})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
