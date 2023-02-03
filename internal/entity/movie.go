package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID              string
	Title           string
	Synopsis        string
	ImdbRating      string
	Votes           int32
	YouChooseRating float32
	Poster          string
	IsDeleted       bool
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       string
	Directors       []*Director
	Actors          []*Actor
	Writers         []*Writer
	Genres          []*Genre
}

func (m *Movie) AddVote() {
	votes := m.GetVotes()
	votes = votes + 1
	m.Votes = votes
}

func (m *Movie) GetVotes() int32 {
	return m.Votes
}

func NewMovie(title string, synopsis string, imdbRating string, poster string) (*Movie, error) {
	dateNow := time.Now()
	movie := &Movie{
		ID:              uuid.New().String(),
		Title:           title,
		Synopsis:        synopsis,
		ImdbRating:      imdbRating,
		Votes:           0,
		YouChooseRating: 0.0,
		Poster:          poster,
		IsDeleted:       false,
		CreatedAt:       dateNow.Local().String(),
		UpdatedAt:       dateNow.Local().String(),
		DeletedAt:       dateNow.Local().String(),
	}

	isValid, err := movie.Validate()
	if !isValid {
		return nil, err
	}

	return movie, nil
}

func (m *Movie) AddDirector(director *Director) {
	m.Directors = append(m.Directors, director)
}

func (m *Movie) RemoveDirector(director *Director) {
	for i, d := range m.Directors {
		if d.ID == director.ID {
			m.Directors = append(m.Directors[:i], m.Directors[i+1:]...)
			return
		}
	}
}

func (m *Movie) AddActor(actor *Actor) {
	m.Actors = append(m.Actors, actor)
}

func (m *Movie) RemoveActor(actor *Actor) {
	for i, d := range m.Actors {
		if d.ID == actor.ID {
			m.Actors = append(m.Actors[:i], m.Actors[i+1:]...)
			return
		}
	}
}

func (m *Movie) AddWriter(writer *Writer) {
	m.Writers = append(m.Writers, writer)
}

func (m *Movie) RemoveWriter(writer *Writer) {
	for i, d := range m.Writers {
		if d.ID == writer.ID {
			m.Writers = append(m.Writers[:i], m.Writers[i+1:]...)
			return
		}
	}
}

func (m *Movie) AddGenre(genre *Genre) {
	m.Genres = append(m.Genres, genre)
}

func (m *Movie) RemoveGenre(genre *Genre) {
	for i, d := range m.Genres {
		if d.ID == genre.ID {
			m.Genres = append(m.Genres[:i], m.Genres[i+1:]...)
			return
		}
	}
}

func (movie *Movie) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["title"] = movie.Title
	inputs["synopsis"] = movie.Synopsis
	inputs["imdbRating"] = movie.ImdbRating
	inputs["poster"] = movie.Poster

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}
