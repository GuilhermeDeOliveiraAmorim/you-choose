package main

import (
	"fmt"
	"time"

	_ "github.com/GuilhermeDeOliveiraAmorim/you-choose/api"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/interface/handlers"
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

	repositories_implementation.Migration(db, sqlDB)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.FRONT_END_URL_VAR.FRONT_END_URL_DEV, config.FRONT_END_URL_VAR.FRONT_END_URL_PROD},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	inputFactory := factories.ImputFactory{
		DB:         db,
		BucketName: config.GOOGLE_VAR.IMAGE_BUCKET_NAME,
	}

	movieFactory := factories.NewMovieFactory(inputFactory)
	listFactory := factories.NewListFactory(inputFactory)
	voteFactory := factories.NewVoteFactory(inputFactory)
	userFactory := factories.NewUserFactory(inputFactory)
	brandFactory := factories.NewBrandFactory(inputFactory)
	middlewareFactory := factories.NewMiddlewareFactory(inputFactory)

	movieHandler := handlers.NewMovieHandler(movieFactory)
	listHandler := handlers.NewListHandler(listFactory)
	voteHandler := handlers.NewVoteHandler(voteFactory)
	userHandler := handlers.NewUserHandler(userFactory)
	brandHandler := handlers.NewBrandHandler(brandFactory)

	public := r.Group("/")
	{
		public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		public.POST("signup", userHandler.CreateUser)
		public.POST("login", userHandler.Login)
		public.GET("lists", listHandler.GetListByID)
		public.GET("lists/all", listHandler.GetLists)
		public.GET("items", listHandler.ShowsRankingItems)
	}

	protectedUser := r.Group("/").Use(middlewareFactory.AuthMiddleware())
	{
		protectedUser.GET("lists/users", listHandler.GetListByUserID)
		protectedUser.POST("votes", voteHandler.Vote)
	}

	protectedAdmin := r.Group("/").Use(middlewareFactory.AuthMiddleware(), middlewareFactory.AdminMiddleware())
	{
		protectedAdmin.POST("lists", listHandler.CreateList)
		protectedAdmin.POST("lists/movies", listHandler.AddMoviesList)
		protectedAdmin.POST("lists/brands", listHandler.AddBrandsList)
		protectedAdmin.POST("items/movies", movieHandler.CreateMovie)
		protectedAdmin.POST("items/brands", brandHandler.CreateBrand)
	}

	r.Run(":8080")
}
