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

func (movie *Movie) AddVote() {
	votes := movie.GetVotes()
	votes = votes + 1
	movie.Votes = votes
}

func (movie *Movie) GetVotes() int32 {
	return movie.Votes
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

func (movie *Movie) AddDirector(director *Director) {
	movie.Directors = append(movie.Directors, director)
}

func (movie *Movie) RemoveDirector(director *Director) {
	for i, d := range movie.Directors {
		if d.ID == director.ID {
			movie.Directors = append(movie.Directors[:i], movie.Directors[i+1:]...)
			return
		}
	}
}

func (movie *Movie) AddActor(actor *Actor) {
	movie.Actors = append(movie.Actors, actor)
}

func (movie *Movie) RemoveActor(actor *Actor) {
	for i, d := range movie.Actors {
		if d.ID == actor.ID {
			movie.Actors = append(movie.Actors[:i], movie.Actors[i+1:]...)
			return
		}
	}
}

func (movie *Movie) AddWriter(writer *Writer) {
	movie.Writers = append(movie.Writers, writer)
}

func (movie *Movie) RemoveWriter(writer *Writer) {
	for i, d := range movie.Writers {
		if d.ID == writer.ID {
			movie.Writers = append(movie.Writers[:i], movie.Writers[i+1:]...)
			return
		}
	}
}

func (movie *Movie) AddGenre(genre *Genre) {
	movie.Genres = append(movie.Genres, genre)
}

func (movie *Movie) RemoveGenre(genre *Genre) {
	for i, d := range movie.Genres {
		if d.ID == genre.ID {
			movie.Genres = append(movie.Genres[:i], movie.Genres[i+1:]...)
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
