package entities

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"

type Vote struct {
	SharedEntity
	ListID        string `json:"list_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
	WinnerID      string `json:"winner_id"`
}

func NewVote(listID, firstMovieID, secondMovieID, winnerID string) (*Vote, []util.ProblemDetails) {
	return &Vote{
		SharedEntity:  *NewSharedEntity(),
		ListID:        listID,
		FirstMovieID:  firstMovieID,
		SecondMovieID: secondMovieID,
		WinnerID:      winnerID,
	}, nil
}
