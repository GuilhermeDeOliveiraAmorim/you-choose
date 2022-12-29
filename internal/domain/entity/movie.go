package entity

import "github.com/google/uuid"

type Movie struct {
	ID              string
	Title           string
	Synopsis        string
	ImdbRating      float32
	Votes           int32
	YouChooseRating float32
	Poster          string
	Directors       []*Director
	Actors          []*Actor
	Writers         []*Writer
	Genres          []*Genre
}

func NewMovie(title string, synopsis string, imdbRating float32, poster string) *Movie {
	return &Movie{
		ID:              uuid.New().String(),
		Title:           title,
		Synopsis:        synopsis,
		ImdbRating:      imdbRating,
		Votes:           0,
		YouChooseRating: 0.0,
		Poster:          poster,
	}
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
