package entities

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type List struct {
	SharedEntity
	Name         string        `json:"name"`
	Movies       []Movie       `json:"movies"`
	Combinations []Combination `json:"combinations"`
}

func NewList(name string) (*List, []util.ProblemDetails) {
	return &List{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
	}, nil
}

func (l *List) AddMovies(movies []Movie) {
	timeNow := time.Now()
	l.Movies = append(l.Movies, movies...)
	l.UpdatedAt = timeNow
}

func (l *List) ClearMovies() {
	l.Movies = []Movie{}
}

func (l *List) AddCombinations(combinations []Combination) {
	l.Combinations = append(l.Combinations, combinations...)
}
