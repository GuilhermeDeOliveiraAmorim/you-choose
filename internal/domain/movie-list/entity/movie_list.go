package domain

import (
	"errors"
	"time"

	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"

	"github.com/google/uuid"
)

type MovieList struct {
	ID          string
	Title       string
	Description string
	Picture     string
	Chooser     []*chooser.Chooser
	Movies      []*movie.Movie
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewMovieList(title string, description string) (*MovieList, error) {
	ml := &MovieList{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
	}

	err := ml.Validate()

	if err != nil {
		return nil, err
	}

	return ml, nil
}

func (ml *MovieList) AddChooser(chooser *chooser.Chooser) {
	ml.Chooser = append(ml.Chooser, chooser)
}

func (ml *MovieList) RemoveChooser(chooser *chooser.Chooser) {
	for i, c := range ml.Chooser {
		if c.ID == chooser.ID {
			ml.Chooser = append(ml.Chooser[:i], ml.Chooser[i+1:]...)
			return
		}
	}
}

func (ml *MovieList) AddMovie(movie *movie.Movie) {
	ml.Movies = append(ml.Movies, movie)
}

func (ml *MovieList) RemoveMovie(movie *movie.Movie) {
	for i, m := range ml.Movies {
		if m.ID == movie.ID {
			ml.Movies = append(ml.Movies[:i], ml.Movies[i+1:]...)
			return
		}
	}
}

func (ml *MovieList) Validate() error {
	if ml.Title == "" || ml.Description == "" || ml.Picture == "" {
		return errors.New("invalid entity")
	}
	return nil
}
