package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/database"
	repositories_implementation "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type VoteFactory struct {
	Vote *usecases.VoteUseCase
}

func NewVoteFactory(input database.StorageInput) *VoteFactory {
	voteResository := repositories_implementation.NewVoteRepository(input.DB)
	listRepository := repositories_implementation.NewListRepository(input.DB)
	movieResository := repositories_implementation.NewMovieRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)
	brandRepository := repositories_implementation.NewBrandRepository(input.DB)

	createVote := usecases.NewVoteUseCase(voteResository, listRepository, movieResository, userResository, brandRepository)

	return &VoteFactory{
		Vote: createVote,
	}
}
