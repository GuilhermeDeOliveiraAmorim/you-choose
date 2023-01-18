package domain

import (
	movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
)

type MovieListRepositoryInterface interface {
	Create(a *movieList.MovieList) (*movieList.MovieList, error)
	Update(a *movieList.MovieList) (*movieList.MovieList, error)
	FindById(id string) (*movieList.MovieList, error)
	DeleteById(id string) (*movieList.MovieList, error)
	FindAll() ([]*movieList.MovieList, error)
}
