package entities

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/oklog/ulid/v2"
)

type Vote struct {
	ID            string     `json:"id"`
	Active        bool       `json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	DeactivatedAt *time.Time `json:"deactivated_at"`
	UserID        string     `json:"user_id"`
	CombinationID string     `json:"combination_id"`
	WinnerID      string     `json:"winner_id"`
}

func NewVote(userID, combinationID, winnerID string) (*Vote, []exceptions.ProblemDetails) {
	return &Vote{
		ID:            ulid.Make().String(),
		Active:        true,
		CreatedAt:     time.Now(),
		DeactivatedAt: nil,
		UserID:        userID,
		CombinationID: combinationID,
		WinnerID:      winnerID,
	}, nil
}
