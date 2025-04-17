package main

import (
	"context"

	_ "github.com/GuilhermeDeOliveiraAmorim/you-choose/api"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/handlers"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/routes"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
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
	ctx := context.Background()

	util.InitLogger()
	util.SetLanguage(config.AVAILABLE_LANGUAGES_VAR.PT_BR)

	db, sqlDB := util.SetupDatabaseConnection(ctx, util.LOCAL)

	models.Migration(ctx, db, sqlDB)

	inputFactory := util.ImputFactory{
		DB:         db,
		BucketName: config.GOOGLE_VAR.IMAGE_BUCKET_NAME,
	}

	handlerFactory := handlers.NewHandlerFactory(inputFactory)
	middlewareFactory := factories.NewMiddlewareFactory(inputFactory)

	router := routes.SetupRouter(handlerFactory, middlewareFactory)

	router.Run(":8080")
}
