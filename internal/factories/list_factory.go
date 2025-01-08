package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type ListFactory struct {
	CreateList      *usecases.CreateListUseCase
	AddMoviesList   *usecases.AddMoviesListUseCase
	GetListByUserID *usecases.GetListByUserIDUseCase
	GetListByID     *usecases.GetListByIDUseCase
	GetLists        *usecases.GetListsUseCase
}

func NewListFactory(input ImputFactory) *ListFactory {
	listRepository := repositories_implementation.NewListRepository(input.DB)
	movieResository := repositories_implementation.NewMovieRepository(input.DB)
	voteRepository := repositories_implementation.NewVoteRepository(input.DB)
	combinationRepository := repositories_implementation.NewCombinationRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)
	imageRepository := repositories_implementation.NewImageRepository(input.BucketName)
	brandRepository := repositories_implementation.NewBrandRepository(input.DB)

	createList := usecases.NewCreateListUseCase(listRepository, movieResository, userResository, imageRepository, brandRepository)
	addMoviesList := usecases.NewAddMoviesListUseCase(listRepository, movieResository, userResository)
	GetListByUserID := usecases.NewGetListByUserIDUseCase(listRepository, voteRepository, combinationRepository, userResository)
	GetListByID := usecases.NewGetListByIDUseCase(listRepository, voteRepository)
	GetLists := usecases.NewGetListsUseCase(listRepository)

	return &ListFactory{
		CreateList:      createList,
		AddMoviesList:   addMoviesList,
		GetListByUserID: GetListByUserID,
		GetListByID:     GetListByID,
		GetLists:        GetLists,
	}
}
