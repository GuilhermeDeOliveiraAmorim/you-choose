package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/middlewares"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-gonic/gin"
)

type MiddlewareFactory struct {
	AuthMiddleware  func() gin.HandlerFunc
	AdminMiddleware func() gin.HandlerFunc
}

func NewMiddlewareFactory(input util.ImputFactory) *MiddlewareFactory {
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
