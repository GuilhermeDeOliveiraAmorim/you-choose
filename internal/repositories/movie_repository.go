package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type MovieRepository interface {
	CreateMovie(movie entities.Movie) error
	GetMovieByID(movieID string) (entities.Movie, error)
	ThisMovieExist(movieExternalID string) (bool, error)
	GetMovies() ([]entities.Movie, error)
	GetMoviesByID(moviesIDs []string) ([]entities.Movie, error)
	SavePoster(poster string) (string, error)
}
