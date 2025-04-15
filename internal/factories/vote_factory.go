package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type VoteFactory struct {
	Vote *usecases.VoteUseCase
}

func NewVoteFactory(input util.ImputFactory) *VoteFactory {
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
