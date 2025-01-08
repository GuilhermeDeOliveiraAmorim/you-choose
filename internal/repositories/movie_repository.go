package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type MovieRepository interface {
	CreateMovie(movie entities.Movie) error
	GetMovieByID(movieID string) (entities.Movie, error)
	ThisMovieExist(movieExternalID string) (bool, error)
	GetMoviesByIDs(moviesIDs []string) ([]entities.Movie, error)
	UpdadeMovie(movie entities.Movie) error
}
