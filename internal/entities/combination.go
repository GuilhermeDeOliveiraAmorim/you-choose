package entities

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
	"github.com/oklog/ulid/v2"
)

type Combination struct {
	ID            string `json:"id"`
	ListID        string `json:"list_id"`
	FirstMovieID  string `json:"first_movie"`
	SecondMovieID string `json:"second_movie"`
}

func NewCombination(listId, firstMovie, secondMovie string) (*Combination, []util.ProblemDetails) {
	return &Combination{
		ID:            ulid.Make().String(),
		ListID:        listId,
		FirstMovieID:  firstMovie,
		SecondMovieID: secondMovie,
	}, nil
}

func (c *Combination) Equals(combination Combination) bool {
	return c.FirstMovieID == combination.FirstMovieID && c.SecondMovieID == combination.SecondMovieID && c.ListID == combination.ListID
}
