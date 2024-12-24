package entities

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"

type Movie struct {
	SharedEntity
	Name   string `json:"name"`
	Year   int64  `json:"year"`
	Poster string `json:"poster"`
}

func NewMovie(name string, year int64, poster string) (*Movie, []util.ProblemDetails) {
	return &Movie{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Year:         year,
		Poster:       poster,
	}, nil
}
