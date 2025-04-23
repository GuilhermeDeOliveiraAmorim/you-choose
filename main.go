package main

import (
	"context"

	_ "github.com/GuilhermeDeOliveiraAmorim/you-choose/api"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/config"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/language"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/logging"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/models"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/routes"
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

	logging.InitLogger()
	language.SetLanguage(config.AVAILABLE_LANGUAGES_VAR.PT_BR)

	db, sqlDB := database.SetupDatabaseConnection(ctx, database.LOCAL)

	models.Migration(ctx, db, sqlDB)

	router := routes.SetupRouter(database.StorageInput{
		DB:         db,
		BucketName: config.GOOGLE_VAR.IMAGE_BUCKET_NAME,
	})

	router.Run(":8080")
}
