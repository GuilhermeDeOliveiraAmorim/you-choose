package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			NewLogger(Logger{
				Code:    401,
				Message: GetErrorMessage("AuthMiddleware", "UnauthorizedHeader").Detail,
				From:    "AuthMiddleware",
				Layer:   "Infra",
				TypeLog: "ERROR",
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": NewProblemDetails(Unauthorized, GetErrorMessage("AuthMiddleware", "UnauthorizedHeader")),
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			NewLogger(Logger{
				Code:    401,
				Message: GetErrorMessage("AuthMiddleware", "UnauthorizedBearer").Detail,
				From:    "AuthMiddleware",
				Layer:   "Infra",
				TypeLog: "ERROR",
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": NewProblemDetails(Unauthorized, GetErrorMessage("AuthMiddleware", "UnauthorizedBearer")),
			})
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				NewLogger(Logger{
					Code:    401,
					Message: GetErrorMessage("AuthMiddleware", "UnauthorizedTokenParse").Detail,
					From:    "AuthMiddleware",
					Layer:   "Infra",
					TypeLog: "ERROR",
				})
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": NewProblemDetails(Unauthorized, GetErrorMessage("AuthMiddleware", "UnauthorizedTokenParse")),
				})
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(config.SECRETS_VAR.JWT_SECRET), nil
		})

		if err != nil {
			NewLogger(Logger{
				Code:    401,
				Message: GetErrorMessage("AuthMiddleware", "UnauthorizedInvalidToken").Detail,
				From:    "AuthMiddleware",
				Layer:   "Infra",
				TypeLog: "ERROR",
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": NewProblemDetails(Unauthorized, GetErrorMessage("AuthMiddleware", "UnauthorizedInvalidToken")),
			})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": NewProblemDetails(Unauthorized, GetErrorMessage("AuthMiddleware", "UnauthorizedToken")),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string)
		c.Set("userID", userID)
		c.Next()
	}
}
