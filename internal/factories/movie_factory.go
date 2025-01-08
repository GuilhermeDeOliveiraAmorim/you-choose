package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type MovieFactory struct {
	CreateMovie *usecases.CreateMovieUseCase
}

func NewMovieFactory(input ImputFactory) *MovieFactory {
	movieResository := repositories_implementation.NewMovieRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)
	imageRepository := repositories_implementation.NewImageRepository(input.BucketName)

	createMovie := usecases.NewCreateMovieUseCase(movieResository, userResository, imageRepository)

	return &MovieFactory{
		CreateMovie: createMovie,
	}
}
