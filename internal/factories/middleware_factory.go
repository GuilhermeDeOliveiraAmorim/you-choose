package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	repositories_implementation "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type MiddlewareFactory struct {
	AuthMiddleware  func() gin.HandlerFunc
	AdminMiddleware func() gin.HandlerFunc
}

func NewMiddlewareFactory(input database.StorageInput) *MiddlewareFactory {
	userRepository := repositories_implementation.NewUserRepository(input.DB)

	return &MiddlewareFactory{
		AuthMiddleware: func() gin.HandlerFunc {
			return middlewares.NewAuthMiddleware(userRepository)
		},
		AdminMiddleware: func() gin.HandlerFunc {
			return middlewares.NewAdminMiddleware(userRepository)
		},
	}
}
