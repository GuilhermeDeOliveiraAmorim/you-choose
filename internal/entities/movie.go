package entities

import (
	"time"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
)

type Movie struct {
	SharedEntity
	Votable
	Name       string `json:"name"`
	Year       int64  `json:"year"`
	Poster     string `json:"poster"`
	ExternalID string `json:"external_id"`
}

func NewMovie(name string, year int64, externalID string) (*Movie, []exceptions.ProblemDetails) {
	return &Movie{
		SharedEntity: *NewSharedEntity(),
		Votable:      *NewVotable(),
		Name:         name,
		Year:         year,
		ExternalID:   externalID,
	}, nil
}

func (m *Movie) AddPoster(poster string) {
	m.Poster = poster
}

func (m *Movie) UpdatePoster(poster string) {
	timeNow := time.Now()
	m.UpdatedAt = &timeNow

	m.Poster = poster
}

func (m *Movie) Equals(movie Movie) bool {
	return m.Name == movie.Name && m.Year == movie.Year && m.ExternalID == movie.ExternalID
}
