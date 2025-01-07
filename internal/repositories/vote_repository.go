package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type VoteRepository interface {
	CreateVote(movie entities.Vote) error
	GetVotesByUserIDAndListID(userID, listID string) ([]entities.Vote, error)
	GetNumberOfVotesByListID(listID string) (int, error)
	VoteAlreadyRegistered(userID, combinationID string) (bool, error)
}
