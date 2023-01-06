package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"

	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
	director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
	genre "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/genre/entity"
	writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"
)

type Movie struct {
	ID              string
	Title           string
	Synopsis        string
	ImdbRating      float32
	Votes           int32
	YouChooseRating float32
	Poster          string
	Directors       []*director.Director
	Actors          []*actor.Actor
	Writers         []*writer.Writer
	Genres          []*genre.Genre
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (m *Movie) AddVote() {
	votes := m.GetVotes()
	votes = votes + 1
	m.Votes = votes
}

func (m *Movie) GetVotes() int32 {
	return m.Votes
}

func NewMovie(title string, synopsis string, imdbRating float32, poster string) (*Movie, error) {
	m := &Movie{
		ID:              uuid.New().String(),
		Title:           title,
		Synopsis:        synopsis,
		ImdbRating:      imdbRating,
		Votes:           0,
		YouChooseRating: 0.0,
		Poster:          poster,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err := m.Validate()

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Movie) AddDirector(director *director.Director) {
	m.Directors = append(m.Directors, director)
}

func (m *Movie) RemoveDirector(director *director.Director) {
	for i, d := range m.Directors {
		if d.ID == director.ID {
			m.Directors = append(m.Directors[:i], m.Directors[i+1:]...)
			return
		}
	}
}

func (m *Movie) AddActor(actor *actor.Actor) {
	m.Actors = append(m.Actors, actor)
}

func (m *Movie) RemoveActor(actor *actor.Actor) {
	for i, d := range m.Actors {
		if d.ID == actor.ID {
			m.Actors = append(m.Actors[:i], m.Actors[i+1:]...)
			return
		}
	}
}

func (m *Movie) AddWriter(writer *writer.Writer) {
	m.Writers = append(m.Writers, writer)
}

func (m *Movie) RemoveWriter(writer *writer.Writer) {
	for i, d := range m.Writers {
		if d.ID == writer.ID {
			m.Writers = append(m.Writers[:i], m.Writers[i+1:]...)
			return
		}
	}
}

func (m *Movie) AddGenre(genre *genre.Genre) {
	m.Genres = append(m.Genres, genre)
}

func (m *Movie) RemoveGenre(genre *genre.Genre) {
	for i, d := range m.Genres {
		if d.ID == genre.ID {
			m.Genres = append(m.Genres[:i], m.Genres[i+1:]...)
			return
		}
	}
}

func (m *Movie) Validate() error {
	if m.Title == "" || m.Synopsis == "" || m.ImdbRating < 0.0 || m.Poster == "" {
		return errors.New("invalid entity")
	}
	return nil
}
