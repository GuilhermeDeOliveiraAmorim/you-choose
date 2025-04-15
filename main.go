package main

import (
	"fmt"
	"time"

	_ "github.com/GuilhermeDeOliveiraAmorim/you-choose/api"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/handlers"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title You Choose API
// @version 1.0
// @description This is an API for managing expenses.
// @termsOfService http://swagger.io/terms/

// @contact.name You Choose
// @contact.url http://www.youchoose.com.br
// @contact.email contato@youchoose.com.br

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	util.SetLanguage(config.AVAILABLE_LANGUAGES_VAR.PT_BR)

	db, sqlDB, err := util.SetupDatabaseConnection(util.LOCAL)
	if err != nil {
		panic("Failed to connect database")
	}
	fmt.Println("Successful connection")

	models.Migration(db, sqlDB)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FRONT_END_URL_VAR.FRONT_END_URL_DEV, config.FRONT_END_URL_VAR.FRONT_END_URL_PROD},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	inputFactory := util.ImputFactory{
		DB:         db,
		BucketName: config.GOOGLE_VAR.IMAGE_BUCKET_NAME,
	}

	handlerFactory := handlers.NewHandlerFactory(inputFactory)
	middlewareFactory := factories.NewMiddlewareFactory(inputFactory)

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

	r.Run(":8080")
}
