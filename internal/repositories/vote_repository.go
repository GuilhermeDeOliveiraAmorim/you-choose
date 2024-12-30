package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type VoteRepository interface {
	CreateVote(movie entities.Vote) error
	GetVotesByUserIDAndListID(userID, listID string) ([]entities.Vote, error)
	VoteAlreadyRegistered(userID, combinationID, listID, winnerID string) (bool, error)
}
