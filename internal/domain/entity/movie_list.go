package entity

import "github.com/google/uuid"

type MovieList struct {
	ID          string
	Title       string
	Description string
	Choosers    []*Chooser
	Movies      []*Movie
}

func NewMovieList(title string, description string) *MovieList {
	return &MovieList{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
	}
}

func (ml *MovieList) AddChooser(chooser *Chooser) {
	ml.Choosers = append(ml.Choosers, chooser)
}

func (ml *MovieList) RemoveChooser(chooser *Chooser) {
	for i, c := range ml.Choosers {
		if c.ID == chooser.ID {
			ml.Choosers = append(ml.Choosers[:i], ml.Choosers[i+1:]...)
			return
		}
	}
}

func (ml *MovieList) AddMovie(movie *Movie) {
	ml.Movies = append(ml.Movies, movie)
}

func (ml *MovieList) RemoveMovie(movie *Movie) {
	for i, m := range ml.Movies {
		if m.ID == movie.ID {
			ml.Movies = append(ml.Movies[:i], ml.Movies[i+1:]...)
			return
		}
	}
}
