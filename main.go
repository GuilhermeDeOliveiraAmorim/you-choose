package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/GuilhermeDeOliveiraAmorim/you-choose/api"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/interface/handlers"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/telemetry"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
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
	shutdown, err := telemetry.InitTracer()
	if err != nil {
		log.Fatalf("Erro ao inicializar o tracer: %v", err)
	}

	defer shutdown(context.Background())

	tracer := otel.Tracer("you-choose-api")

	util.SetLanguage(config.AVAILABLE_LANGUAGES_VAR.ZH_CN)

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
	movieHandler := handlers.NewMovieHandler(movieFactory)

	listFactory := factories.NewListFactory(inputFactory)
	listHandler := handlers.NewListHandler(listFactory)

	voteFactory := factories.NewVoteFactory(inputFactory)
	voteHandler := handlers.NewVoteHandler(voteFactory)

	userFactory := factories.NewUserFactory(inputFactory)
	userHandler := handlers.NewUserHandler(userFactory, tracer)

	brandFactory := factories.NewBrandFactory(inputFactory)
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

	protected := r.Group("/").Use(util.AuthMiddleware())
	{
		protected.POST("lists", listHandler.CreateList)
		protected.POST("lists/movies", listHandler.AddMoviesList)
		protected.POST("lists/brands", listHandler.AddBrandsList)
		protected.GET("lists/users", listHandler.GetListByUserID)

		protected.POST("items/movies", movieHandler.CreateMovie)
		protected.POST("items/brands", brandHandler.CreateBrand)

		protected.POST("votes", voteHandler.Vote)
	}

	r.Run(":8080")
}
