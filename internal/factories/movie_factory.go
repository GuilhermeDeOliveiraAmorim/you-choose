package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type MovieFactory struct {
	CreateMovie *usecases.CreateMovieUseCase
}

func NewMovieFactory(db *gorm.DB, bucketName string) *MovieFactory {
	movieResository := repositories_implementation.NewMovieRepository(db, bucketName)
	userResository := repositories_implementation.NewUserRepository(db)
	imageRepository := repositories_implementation.NewImageRepository(bucketName)

	createMovie := usecases.NewCreateMovieUseCase(movieResository, userResository, imageRepository)

	return &MovieFactory{
		CreateMovie: createMovie,
	}
}
