package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type ListFactory struct {
	CreateList      *usecases.CreateListUseCase
	AddMoviesList   *usecases.AddMoviesListUseCase
	GetListByUserID *usecases.GetListByUserIDUseCase
	GetListByID     *usecases.GetListByIDUseCase
}

func NewListFactory(db *gorm.DB, bucketName string) *ListFactory {
	listRepository := repositories_implementation.NewListRepository(db)
	movieResository := repositories_implementation.NewMovieRepository(db, bucketName)
	voteRepository := repositories_implementation.NewVoteRepository(db)
	combinationRepository := repositories_implementation.NewCombinationRepository(db)
	userResository := repositories_implementation.NewUserRepository(db)

	createList := usecases.NewCreateListUseCase(listRepository, movieResository, userResository)
	addMoviesList := usecases.NewAddMoviesListUseCase(listRepository, movieResository, userResository)
	GetListByUserID := usecases.NewGetListByUserIDUseCase(listRepository, voteRepository, combinationRepository, userResository)
	GetListByID := usecases.NewGetListByIDUseCase(listRepository, voteRepository)

	return &ListFactory{
		CreateList:      createList,
		AddMoviesList:   addMoviesList,
		GetListByUserID: GetListByUserID,
		GetListByID:     GetListByID,
	}
}
