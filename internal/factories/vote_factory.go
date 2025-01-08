package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type VoteFactory struct {
	Vote *usecases.VoteUseCase
}

func NewVoteFactory(input ImputFactory) *VoteFactory {
	voteResository := repositories_implementation.NewVoteRepository(input.DB)
	listRepository := repositories_implementation.NewListRepository(input.DB)
	movieResository := repositories_implementation.NewMovieRepository(input.DB)
	userResository := repositories_implementation.NewUserRepository(input.DB)

	createVote := usecases.NewVoteUseCase(voteResository, listRepository, movieResository, userResository)

	return &VoteFactory{
		Vote: createVote,
	}
}
