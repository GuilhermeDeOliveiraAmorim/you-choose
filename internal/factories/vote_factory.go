package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type VoteFactory struct {
	Vote *usecases.VoteUseCase
}

func NewVoteFactory(db *gorm.DB) *VoteFactory {
	voteResository := repositories_implementation.NewVoteRepository(db)
	listRepository := repositories_implementation.NewListRepository(db)
	movieResository := repositories_implementation.NewMovieRepository(db)

	createVote := usecases.NewVoteUseCase(voteResository, listRepository, movieResository)

	return &VoteFactory{
		Vote: createVote,
	}
}
