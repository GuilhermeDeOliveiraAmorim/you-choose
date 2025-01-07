package factories

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infrastructure/repositories_implementation"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	"gorm.io/gorm"
)

type VoteFactory struct {
	Vote *usecases.VoteUseCase
}

func NewVoteFactory(db *gorm.DB, bucketName string) *VoteFactory {
	voteResository := repositories_implementation.NewVoteRepository(db)
	listRepository := repositories_implementation.NewListRepository(db)
	movieResository := repositories_implementation.NewMovieRepository(db, bucketName)
	userResository := repositories_implementation.NewUserRepository(db)

	createVote := usecases.NewVoteUseCase(voteResository, listRepository, movieResository, userResository)

	return &VoteFactory{
		Vote: createVote,
	}
}
