package entities

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type List struct {
	SharedEntity
	Name         string        `json:"name"`
	Movies       []Movie       `json:"movies"`
	Combinations []Combination `json:"combinations"`
}

func NewList(name string) (*List, []util.ProblemDetails) {
	return &List{
		SharedEntity: *NewSharedEntity(),
		Name:         name,
	}, nil
}

func (l *List) AddMovies(movies []Movie) {
	if len(l.Movies) == 0 {
		l.Movies = movies
		return
	}

	uniqueMovies := []Movie{}
	for _, newMovie := range movies {
		exists := false
		for _, existingMovie := range l.Movies {
			if existingMovie.Equals(newMovie) {
				exists = true
				break
			}
		}
		if !exists {
			uniqueMovies = append(uniqueMovies, newMovie)
		}
	}

	l.Movies = uniqueMovies
}

func (l *List) ClearMovies() {
	l.Movies = []Movie{}
}

func (l *List) AddCombinations(combinations []Combination) {
	if len(l.Combinations) == 0 {
		l.Combinations = combinations
		return
	}

	uniqueCombinations := []Combination{}
	for _, newCombination := range combinations {
		exists := false
		for _, existingCombination := range l.Combinations {
			if existingCombination.Equals(newCombination) {
				exists = true
				break
			}
		}
		if !exists {
			uniqueCombinations = append(uniqueCombinations, newCombination)
		}
	}

	l.Combinations = uniqueCombinations
}

func (l *List) GetCombinations(movies []string) ([]Combination, []util.ProblemDetails) {
	var combinations []Combination

	for i := 0; i < len(movies); i++ {
		for j := i + 1; j < len(movies); j++ {
			newCombination, errNewCombination := NewCombination(l.ID, movies[i], movies[j])
			if errNewCombination != nil {
				return []Combination{}, errNewCombination
			}

			combinations = append(combinations, *newCombination)
		}
	}

	return combinations, nil
}

func (l *List) GetMovieIDs() ([]string, []util.ProblemDetails) {
	movieIDs := []string{}

	for _, movie := range l.Movies {
		movieIDs = append(movieIDs, movie.ID)
	}

	return movieIDs, nil
}
