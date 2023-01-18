package domain

import (
	movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"
)

type MovieRepositoryInterface interface {
	UpdateYouChooseRating(id string) (*movie.Movie, error)
	Create(a *movie.Movie) (*movie.Movie, error)
	Update(a *movie.Movie) (*movie.Movie, error)
	DeleteById(id string) (*movie.Movie, error)
	FindById(id string) (*movie.Movie, error)
	FindAll() ([]*movie.Movie, error)
}
