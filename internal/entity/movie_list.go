package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type MovieList struct {
	ID          string
	Title       string
	Description string
	Picture     string
	Choosers    []*Chooser
	Movies      []*Movie
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	IsDeleted   bool
}

func NewMovieList(title string, description string, picture string) (*MovieList, error) {
	movieList := &MovieList{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Picture:     picture,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
		IsDeleted:   false,
	}

	isValidMovieList, err := movieList.Validate()
	if !isValidMovieList {
		return nil, err
	}

	return movieList, nil
}

func (movieList *MovieList) AddChooser(chooser *Chooser) {
	movieList.Choosers = append(movieList.Choosers, chooser)
}

func (movieList *MovieList) RemoveChooser(chooserForRemove *Chooser) {
	for position, chooser := range movieList.Choosers {
		if chooser.ID == chooserForRemove.ID {
			movieList.Choosers = append(movieList.Choosers[:position], movieList.Choosers[position+1:]...)
			return
		}
	}
}

func (movieList *MovieList) AddMovie(movie *Movie) {
	movieList.Movies = append(movieList.Movies, movie)
}

func (movieList *MovieList) RemoveMovie(movieForRemove *Movie) {
	for position, movie := range movieList.Movies {
		if movie.ID == movieForRemove.ID {
			movieList.Movies = append(movieList.Movies[:position], movieList.Movies[position+1:]...)
			return
		}
	}
}

func (movieList *MovieList) Validate() (bool, error) {
	inputs := make(map[string]string)

	inputs["title"] = movieList.Title
	inputs["description"] = movieList.Description
	inputs["picture"] = movieList.Picture

	for key, value := range inputs {
		if value == "" {
			message := key + " cannot be empty"
			return false, errors.New(message)
		}
	}

	return true, nil
}