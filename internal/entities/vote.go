package entities

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Vote struct {
	SharedEntity
	UserID        string `json:"user_id"`
	CombinationID string `json:"combination_id"`
	ListID        string `json:"list_id"`
	WinnerID      string `json:"winner_id"`
}

func NewVote(userID, combinationID, listID, winnerID string) (*Vote, []util.ProblemDetails) {
	return &Vote{
		SharedEntity:  *NewSharedEntity(),
		UserID:        userID,
		CombinationID: combinationID,
		ListID:        listID,
		WinnerID:      winnerID,
	}, nil
}
