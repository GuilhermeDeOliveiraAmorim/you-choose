package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type MovieFactory struct {
	CreateMovie *usecases.CreateMovieUseCase
}

func NewMovieFactory(db *gorm.DB) *MovieFactory {
	movieResository := repositories_implementation.NewMovieRepository(db)

	createMovie := usecases.NewCreateMovieUseCase(movieResository)

	return &MovieFactory{
		CreateMovie: createMovie,
	}
}
