package factories

import (
	repositories_implementation "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type ListFactory struct {
	CreateList        *usecases.CreateListUseCase
	AddMoviesList     *usecases.AddMoviesListUseCase
	GetListByUserID   *usecases.GetListByUserIDUseCase
	GetListByID       *usecases.GetListByIDUseCase
	GetLists          *usecases.GetListsUseCase
	AddBrandsList     *usecases.AddBrandsListUseCase
	ShowsRankingItems *usecases.ShowsRankingItemsUseCase
}

func NewListFactory(input util.ImputFactory) *ListFactory {
	listRepository := repositories_implementation.NewListRepository(input.DB)
	movieResository := repositories_implementation.NewMovieRepository(input.DB)
	voteRepository := repositories_implementation.NewVoteRepository(input.DB)
	combinationRepository := repositories_implementation.NewCombinationRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)
	imageRepository := repositories_implementation.NewImageRepository(input.BucketName)
	brandRepository := repositories_implementation.NewBrandRepository(input.DB)

	createList := usecases.NewCreateListUseCase(listRepository, movieResository, userResository, imageRepository, brandRepository)
	addMoviesList := usecases.NewAddMoviesListUseCase(listRepository, movieResository, userResository)
	getListByUserID := usecases.NewGetListByUserIDUseCase(listRepository, voteRepository, combinationRepository, userResository)
	getListByID := usecases.NewGetListByIDUseCase(listRepository, voteRepository)
	getLists := usecases.NewGetListsUseCase(listRepository)
	addBrandsList := usecases.NewAddBrandsListUseCase(listRepository, brandRepository, userResository)
	showsRankingItems := usecases.NewShowsRankingItemsUseCase(movieResository, brandRepository)

	return &ListFactory{
		CreateList:        createList,
		AddMoviesList:     addMoviesList,
		GetListByUserID:   getListByUserID,
		GetListByID:       getListByID,
		GetLists:          getLists,
		AddBrandsList:     addBrandsList,
		ShowsRankingItems: showsRankingItems,
	}
}
