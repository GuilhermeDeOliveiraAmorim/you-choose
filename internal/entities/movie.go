package entities

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"

type Movie struct {
	SharedEntity
	Name       string `json:"name"`
	Year       int64  `json:"year"`
	Poster     string `json:"poster"`
	ExternalID string `json:"external_id"`
}

func NewMovie(name string, year int64, externalID string) (*Movie, []util.ProblemDetails) {
	return &Movie{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
		Year:         year,
		ExternalID:   externalID,
	}, nil
}

func (m *Movie) UpdatePoster(poster string) {
	m.Poster = poster
}

func (m *Movie) Equals(movie Movie) bool {
	return m.Name == movie.Name && m.Year == movie.Year && m.ExternalID == movie.ExternalID
}
