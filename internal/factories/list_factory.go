package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type ListFactory struct {
	CreateList    *usecases.CreateListUseCase
	AddMoviesList *usecases.AddMoviesListUseCase
	GetListByID   *usecases.GetListByIDUseCase
}

func NewListFactory(db *gorm.DB) *ListFactory {
	listRepository := repositories_implementation.NewListRepository(db)
	movieResository := repositories_implementation.NewMovieRepository(db)
	voteRepository := repositories_implementation.NewVoteRepository(db)
	combinationRepository := repositories_implementation.NewCombinationRepository(db)

	createList := usecases.NewCreateListUseCase(listRepository, movieResository)
	addMoviesList := usecases.NewAddMoviesListUseCase(listRepository, movieResository)
	getListByID := usecases.NewGetListByIDUseCase(listRepository, voteRepository, combinationRepository)

	return &ListFactory{
		CreateList:    createList,
		AddMoviesList: addMoviesList,
		GetListByID:   getListByID,
	}
}
