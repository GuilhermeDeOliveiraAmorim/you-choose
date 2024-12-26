package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type ListFactory struct {
	CreateList    *usecases.CreateListUseCase
	AddMoviesList *usecases.AddMoviesListUseCase
}

func NewCategoryFactory(db *gorm.DB) *ListFactory {
	listRepository := repositories_implementation.NewListRepository(db)
	movieResository := repositories_implementation.NewMovieRepository(db)

	createList := usecases.NewCreateListUseCase(listRepository)
	addMoviesList := usecases.NewAddMoviesListUseCase(listRepository, movieResository)

	return &ListFactory{
		CreateList:    createList,
		AddMoviesList: addMoviesList,
	}
}
