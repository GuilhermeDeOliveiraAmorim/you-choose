package routes

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(storageInput database.StorageInput) *gin.Engine {
	handlerFactory := handlers.NewHandlerFactory(storageInput)
	middlewareFactory := factories.NewMiddlewareFactory(storageInput)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FRONT_END_URL_VAR.FRONT_END_URL_DEV, config.FRONT_END_URL_VAR.FRONT_END_URL_PROD},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	public := r.Group("/")
	{
		public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		public.POST("signup", handlerFactory.UserHandler.CreateUser)
		public.POST("login", handlerFactory.UserHandler.Login)
		public.GET("lists", handlerFactory.ListHandler.GetListByID)
		public.GET("lists/all", handlerFactory.ListHandler.GetLists)
		public.GET("items", handlerFactory.ListHandler.ShowsRankingItems)
	}

	protectedUser := r.Group("/").Use(middlewareFactory.AuthMiddleware())
	{
		protectedUser.GET("lists/users", handlerFactory.ListHandler.GetListByUserID)
		protectedUser.POST("votes", handlerFactory.VoteHandler.Vote)
	}

	protectedAdmin := r.Group("/").Use(middlewareFactory.AuthMiddleware(), middlewareFactory.AdminMiddleware())
	{
		protectedAdmin.POST("lists", handlerFactory.ListHandler.CreateList)
		protectedAdmin.POST("lists/movies", handlerFactory.ListHandler.AddMoviesList)
		protectedAdmin.POST("lists/brands", handlerFactory.ListHandler.AddBrandsList)
		protectedAdmin.POST("items/movies", handlerFactory.MovieHandler.CreateMovie)
		protectedAdmin.POST("items/brands", handlerFactory.BrandHandler.CreateBrand)
	}

	return r
}
